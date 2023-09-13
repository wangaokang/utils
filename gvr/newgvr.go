package gvr

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

func KindToResource(kind string) string {
	kindLower := strings.ToLower(kind)
	if strings.HasSuffix(kindLower, "s") {
		return kindLower
	} else {
		return kindLower + "s"
	}
}

func NewGvr(unstructuredObj unstructured.Unstructured) *schema.GroupVersionResource {
	return &schema.GroupVersionResource{
		Group:    unstructuredObj.GroupVersionKind().Group,
		Version:  unstructuredObj.GroupVersionKind().Version,
		Resource: KindToResource(unstructuredObj.GetKind()),
	}

}
