package envoy

import (
	"reflect"
	"testing"

	envoy_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	networkingv1 "k8s.io/api/networking/v1"
)

func notIsType(a, b interface{}) bool {
	return reflect.TypeOf(a) != reflect.TypeOf(b)
}

func notIsPath(a, b string) bool {
	return a != b
}

func TestPrefixStringPath(t *testing.T) {
	ingress := newIngress("foo.app.com", "foo.cluster.com", pathvars{"path": "/", "pathType": networkingv1.PathTypePrefix})
	c := translateIngresses([]networkingv1.Ingress{ingress})

	pathtype := c.VirtualHosts[0].Routes[0].Route.GetPathSpecifier()
	pathroute := c.VirtualHosts[0].Routes[0].Route.GetPrefix()

	if notIsType(pathtype, &envoy_route_v3.RouteMatch_Prefix{}) {
		t.Errorf("expected PathType for foo.app.com ingress , was %s", reflect.TypeOf(pathtype))
	}
	if notIsPath(pathroute, ingress.Spec.Rules[0].HTTP.Paths[0].Path) {
		t.Errorf("expected PathRoute for foo.app.com ingress , was %s", ingress.Spec.Rules[0].HTTP.Paths[0].Path)
	}
}

func TestPrefixPath(t *testing.T) {
	ingress := newIngress("foo.app.com", "foo.cluster.com", pathvars{"path": "/testing", "pathType": networkingv1.PathTypePrefix})
	c := translateIngresses([]networkingv1.Ingress{ingress})

	pathtype := c.VirtualHosts[0].Routes[0].Route.GetPathSpecifier()
	pathroute := c.VirtualHosts[0].Routes[0].Route.GetPathSeparatedPrefix()

	if notIsType(pathtype, &envoy_route_v3.RouteMatch_PathSeparatedPrefix{}) {
		t.Errorf("expected PathType for foo.app.com ingress , was %s", reflect.TypeOf(pathtype))
	}
	if notIsPath(pathroute, ingress.Spec.Rules[0].HTTP.Paths[0].Path) {
		t.Errorf("expected PathRoute for foo.app.com ingress , was %s", ingress.Spec.Rules[0].HTTP.Paths[0].Path)
	}
}

func TestRegexPath(t *testing.T) {
	ingress := newIngress("foo.app.com", "foo.cluster.com", pathvars{"path": "/foo/.*", "pathType": networkingv1.PathTypeImplementationSpecific})
	c := translateIngresses([]networkingv1.Ingress{ingress})

	pathtype := c.VirtualHosts[0].Routes[0].Route.GetPathSpecifier()
	pathroute := c.VirtualHosts[0].Routes[0].Route.GetSafeRegex().GetRegex()

	if notIsType(pathtype, &envoy_route_v3.RouteMatch_SafeRegex{}) {
		t.Errorf("expected PathType for foo.app.com ingress , was %s", reflect.TypeOf(pathtype))
	}
	if notIsPath(pathroute, "^/foo/.*") {
		t.Errorf("expected PathRoute for foo.app.com ingress , was %s", ingress.Spec.Rules[0].HTTP.Paths[0].Path)
	}
}

func TestExactPath(t *testing.T) {
	ingress := newIngress("foo.app.com", "foo.cluster.com", pathvars{"path": "/foo", "pathType": networkingv1.PathTypeExact})
	c := translateIngresses([]networkingv1.Ingress{ingress})

	pathtype := c.VirtualHosts[0].Routes[0].Route.GetPathSpecifier()
	pathroute := c.VirtualHosts[0].Routes[0].Route.GetPath()

	if notIsType(pathtype, &envoy_route_v3.RouteMatch_Path{}) {
		t.Errorf("expected PathType for foo.app.com ingress , was %s", reflect.TypeOf(pathtype))
	}
	if notIsPath(pathroute, ingress.Spec.Rules[0].HTTP.Paths[0].Path) {
		t.Errorf("expected PathRoute for foo.app.com ingress , was %s", ingress.Spec.Rules[0].HTTP.Paths[0].Path)
	}
}
