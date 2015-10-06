package romulus

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/meta"
	unvApi "k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
)

var (
	bckndsKeyFmt   = "%s/backends"
	frntndsKeyFmt  = "%s/frontends"
	vulcanKeyLabel = "romulus/vulcanKey"

	KubeRetryLimit = 10 * time.Second

	serverTagLen = 8
)

// EtcdPeerList is just a slice of etcd peers
type EtcdPeerList []string

// KubeClientConfig is an alias for kubernetes/pkg/client.Config
type KubeClientConfig unversioned.Config

// ServiceSelector is a map of labels for selecting services
type ServiceSelector map[string]string

func (s ServiceSelector) fixNamespace() ServiceSelector {
	ss := make(map[string]string, len(s))
	for k := range s {
		key := k
		if !strings.HasPrefix(k, "romulus/") {
			key = fmt.Sprintf("romulus/%s", key)
		}
		ss[key] = s[k]
	}
	return ServiceSelector(ss)
}

func formatEtcdNamespace(v string) string {
	return fmt.Sprintf("/%s", strings.Trim(v, "/"))
}

// Config is used to configure the Registrar
type Config struct {
	PeerList            EtcdPeerList
	EtcdTimeout         time.Duration
	KubeConfig          KubeClientConfig
	APIVersion          string
	Selector            ServiceSelector
	VulcanEtcdNamespace string
}

func (c *Config) kc() unversioned.Config { return (unversioned.Config)(c.KubeConfig) }
func (c *Config) ps() []string           { return ([]string)(c.PeerList) }

func (sl ServiceSelector) String() string {
	s := []string{}
	for k, v := range sl {
		s = append(s, strings.Join([]string{k, v}, "="))
	}
	return strings.Join(s, ", ")
}

// Registrar holds the kubernetes/pkg/client.Client and etcd.Client
type Registrar struct {
	k  unversioned.Interface
	e  EtcdClient
	vk string
	v  string
	s  ServiceSelector
}

// NewRegistrar returns a ptr to a new Registrar from a Config
func NewRegistrar(c *Config) (*Registrar, error) {
	cf := c.kc()
	kc, err := unversioned.New(&cf)
	if err != nil {
		return nil, err
	}
	ec, err := NewEtcdClient(c.ps(), formatEtcdNamespace(c.VulcanEtcdNamespace), c.EtcdTimeout)
	if err != nil {
		return nil, err
	}
	return &Registrar{
		e:  ec,
		k:  kc,
		v:  c.APIVersion,
		s:  c.Selector.fixNamespace(),
		vk: formatEtcdNamespace(c.VulcanEtcdNamespace),
	}, nil
}

func (r *Registrar) setVulcanKey(o runtime.Object) {
	m := meta.NewAccessor()
	la, er := m.Labels(o)
	if er != nil {
		r.e.SetPrefix(r.vk)
		return
	}
	if val, ok := la[vulcanKeyLabel]; ok {
		r.e.SetPrefix(formatEtcdNamespace(val))
		return
	}
	r.e.SetPrefix(r.vk)
}

func (r *Registrar) getEndpoint(name, ns string) (en *api.Endpoints, er error) {
	if ns == "" {
		return nil, NewKubeNotFound("Endpoints", name)
	}

	fn := func() error {
		en, er = r.k.Endpoints(ns).Get(name)
		if er == nil || kubeIsNotFound(er) {
			return nil
		}
		return er
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = KubeRetryLimit
	if e := backoff.Retry(fn, b); e != nil {
		return nil, NewErr(e, "kubernetes error")
	}
	return
}

func (r *Registrar) getService(name, ns string) (s *api.Service, er error) {
	if ns == "" {
		return nil, NewKubeNotFound("Service", name)
	}

	fn := func() error {
		s, er = r.k.Services(ns).Get(name)
		if er == nil || kubeIsNotFound(er) {
			return nil
		}
		return er
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = KubeRetryLimit
	if e := backoff.Retry(fn, b); e != nil {
		return nil, NewErr(e, "kubernetes error")
	}
	return
}

func (r *Registrar) pruneServers(bid string, sm ServerMap) error {
	k := fmt.Sprintf(srvrDirFmt, bid)
	srvs, e := r.e.Keys(k)
	if e != nil {
		if isKeyNotFound(e) {
			return nil
		}
		return NewErr(e, "etcd error")
	}

	logf(fi{"servers": sm, "backend": bid}).Debug("Gathered known servers from kubernetes")
	logf(fi{"servers": srvs, "backend": bid}).Debug("Gathered known servers from etcd")
	for _, id := range srvs {
		key := fmt.Sprintf("%s/%s", k, id)
		s, e := r.e.Val(key)
		if e != nil {
			logf(fi{"server": id, "backend": bid}).Warnf("Error getting server from etcd: %v", e)
			continue
		}

		srv := &Server{ID: id, Backend: bid}
		if e := decode(srv, []byte(s)); e != nil {
			logf(srv).Errorf("Unable to unmarshall Server: %v", e)
			logf(srv).Debugf("Data: %s", s)
			if e := r.e.Del(key); e != nil {
				logf(srv).Errorf("Error removing server: %v", e)
				continue
			}
		}

		sTag := md5Hash(bid, srv.URL.String())[:serverTagLen]
		if nSrv, ok := sm[sTag]; ok {
			logf(srv, nSrv).Debug("Server exists, syncing IDs")
			nSrv.ID = srv.ID
		} else {
			logf(srv).Info("Removing Server")
			if e := r.e.Del(key); e != nil {
				logf(srv).Errorf("Error removing server: %v", e)
				continue
			}
		}
	}
	return nil
}

func (r *Registrar) pruneBackends() error {
	ids, err := r.e.Keys(bcknds)
	if err != nil {
		if isKeyNotFound(err) {
			return nil
		}
		return NewErr(err, "etcd error")
	}

	log().Debugf("Found current backends: %v", ids)
	for _, id := range ids {
		name, ns, e := parseVulcanID(id)
		if e != nil {
			logf(fi{"id": id}).Error("Invalid ID")
			key := fmt.Sprintf(bckndDirFmt, id)
			if e := r.e.Del(key); e != nil {
				logf(fi{"backend": id}).Warn("etcd error")
			}
		} else if _, err := r.getEndpoint(name, ns); err != nil && kubeIsNotFound(err) {
			logf(fi{"id": id, "service": name, "namespace": ns}).Warnf("Did not find backend on API server: %v", err)
			b := NewBackend(id)
			if err := r.e.Del(b.DirKey()); err != nil {
				logf(fi{"backend": id}).Warn("etcd error")
			}
		}
	}
	return nil
}

func (r *Registrar) pruneFrontends() error {
	ids, err := r.e.Keys(frntnds)
	if err != nil {
		if isKeyNotFound(err) {
			return nil
		}
		return NewErr(err, "etcd error")
	}

	log().Debugf("Found current frontends: %v", ids)
	for _, id := range ids {
		name, ns, e := parseVulcanID(id)
		if e != nil {
			logf(fi{"id": id}).Error("Invalid ID")
			key := fmt.Sprintf(frntndDirFmt, id)
			if e := r.e.Del(key); e != nil {
				logf(fi{"frontend": id}).Warn("etcd error")
			}
		} else if _, err := r.getService(name, ns); err != nil && kubeIsNotFound(err) {
			logf(fi{"id": id, "service": name, "namespace": ns}).Warnf("Did not find frontend on API server: %v", err)
			f := NewFrontend(id, "")
			if err := r.e.Del(f.DirKey()); err != nil {
				logf(fi{"frontend": id}).Warn("etcd error")
			}
		}
	}
	return nil
}

func (reg *Registrar) delete(r runtime.Object) error {
	reg.setVulcanKey(r)
	switch o := r.(type) {
	case *api.Endpoints:
		return deregisterEndpoints(reg, o)
	case *api.Service:
		return deregisterService(reg, o)
	default:
		return NewErr(nil, "Unsupported api object: %v", r)
	}
}

func (reg *Registrar) update(r runtime.Object, s string) error {
	reg.setVulcanKey(r)
	switch o := r.(type) {
	case *api.Service:
		if s == "mod" {
			return registerService(reg, o)
		}
		return nil
	case *api.Endpoints:
		return registerEndpoint(reg, o)
	default:
		return NewErr(nil, "Unsupported api object: %v", r)
	}
}

func (r *Registrar) registerBackends(s *api.Service, e *api.Endpoints) (*BackendList, error) {
	bnds := NewBackendList()
	logf(fi{"service": e.Name, "namespace": e.Namespace}).Info("Registering backend")
	r.pruneBackends()
	for _, es := range e.Subsets {
		for _, port := range es.Ports {
			if port.Protocol != api.ProtocolTCP {
				logf(fi{"service": e.Name, "namespace": e.Namespace}).Warnf("Unsupported protocol: %s", port.Protocol)
				continue
			}

			sm := ServerMap{}
			bid := getVulcanID(e.Name, e.Namespace, port.Name)
			bnd := NewBackend(bid)

			if st, ok := s.Annotations[bckndSettingsAnnotation]; ok {
				bnd.Settings = NewBackendSettings([]byte(st))
			}
			logf(fi{"id": bnd.ID, "type": bnd.Type, "settings": bnd.Settings.String()}).Debug("Backend settings")

			logf(fi{"addresses": es.Addresses}).Debug("Gathering kubernetes endpoints")
			for _, ip := range es.Addresses {
				ur := fmt.Sprintf("http://%s:%d", ip.IP, port.Port)
				u, err := url.Parse(ur)
				if err != nil {
					logf(fi{"service": e.Name, "namespace": e.Namespace, "id": bnd.ID}).Warnf("Bad URL: %s", ur)
					continue
				}
				uu := (*URL)(u)
				sTag := md5Hash(bid, uu.String())[:serverTagLen]
				sm[sTag] = &Server{
					ID:      fmt.Sprintf("%s-%s", bid, sTag),
					Backend: bid,
					URL:     uu,
				}
			}
			if err := r.pruneServers(bid, sm); err != nil {
				return bnds, NewErr(err, "Unable to prune servers for backend %q", bid)
			}

			val, err := bnd.Val()
			if err != nil {
				return bnds, NewErr(err, "Could not encode backend for %q", e.Name)
			}
			eVal, _ := r.e.Val(bnd.Key())
			if val != eVal {
				logf(bnd).Debug("Upserting backend")
				if err := r.e.Add(bnd.Key(), val); err != nil {
					return bnds, NewErr(err, "etcd error")
				}
			} else {
				logf(fi{"existing": eVal, "new": val}).Debugf("No changes, not upserting Backend %q", bnd.ID)
			}
			bnds.Add(port.Port, port.Name, bnd)

			for _, srv := range sm {
				val, err := srv.Val()
				if err != nil {
					logf(srv).Warnf("Unable to encode server: %v", err)
					continue
				}
				eVal, _ := r.e.Val(srv.Key())
				if val != eVal {
					logf(srv).Debug("Upserting server")
					if err := r.e.Add(srv.Key(), val); err != nil {
						return bnds, NewErr(err, "etcd error")
					}
				} else {
					logf(fi{"existing": eVal, "new": val}).Debugf("No changes, not upserting Server %q", srv.ID)
				}
			}
		}
	}
	return bnds, nil
}

func (r *Registrar) registerFrontends(s *api.Service, bnds *BackendList) error {
	logf(fi{"service": s.Name, "namespace": s.Namespace}).Info("Registering frontend")
	r.pruneFrontends()
	logf(fi{"service": s.Name, "namespace": s.Namespace}).Debugf("Backend List: %+v", bnds)
	for _, port := range s.Spec.Ports {
		bnd, ok := bnds.Lookup(port.TargetPort.IntVal, port.TargetPort.StrVal)
		if !ok {
			logf(fi{"service": s.Name, "namespace": s.Namespace}).
				Warnf("No backend for service port %d (target: %d)", port.Port, port.TargetPort.IntVal)
			continue
		}

		fid := getVulcanID(s.Name, s.Namespace, port.Name)
		fnd := NewFrontend(fid, bnd.ID)
		fnd.Route = buildRoute(port.Name, s.Annotations)
		if st, ok := s.Annotations[frntndSettingsAnnotation]; ok {
			fnd.Settings = NewFrontendSettings([]byte(st))
		}
		logf(fnd).Debug("Frontend settings")

		val, err := fnd.Val()
		if err != nil {
			return NewErr(err, "Could not encode frontend for %q", s.Name)
		}
		eVal, _ := r.e.Val(fnd.Key())
		if val != eVal {
			logf(fi{"id": fnd.ID, "backend": bnd.ID}).Debug("Upserting frontend")
			if err := r.e.Add(fnd.Key(), val); err != nil {
				return NewErr(err, "etcd error")
			}
		} else {
			logf(fi{"existing": eVal, "new": val}).Debugf("No changes, not upserting Frontend %q", fnd.ID)
		}
	}
	return nil
}

func getVulcanID(name, ns, port string) string {
	var id []string
	if port != "" {
		id = []string{port, name, ns}
	} else {
		id = []string{name, ns}
	}
	return strings.Join(id, ".")
}

func parseVulcanID(id string) (string, string, error) {
	bits := strings.Split(id, ".")
	if len(bits) < 2 {
		return "", "", NewErr(nil, "Invalid vulcan ID %q", id)
	}
	return bits[len(bits)-2], bits[len(bits)-1], nil
}

func registerable(o runtime.Object, sl ServiceSelector) bool {
	if _, ok := o.(*unvApi.Status); ok {
		return true
	}
	m := meta.NewAccessor()
	la, er := m.Labels(o)
	if er != nil {
		log().Debugf("Failed to access labels: %v", er)
		return false
	}

	for k, v := range sl.fixNamespace() {
		if val, ok := la[k]; !ok || val != v {
			return false
		}
	}
	return true
}
