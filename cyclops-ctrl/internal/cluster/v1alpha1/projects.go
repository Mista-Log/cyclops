package v1alpha1

import (
	"context"
	"github.com/cyclops-ui/cycops-ctrl/internal/models/crd/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"time"
)

type ModuleInterface interface {
	List(opts metav1.ListOptions) ([]v1alpha1.Module, error)
	Get(name string) (*v1alpha1.Module, error)
	Create(*v1alpha1.Module) (*v1alpha1.Module, error)
	Update(*v1alpha1.Module) (*v1alpha1.Module, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Delete(name string) error
	// ...
}

type moduleClient struct {
	restClient rest.Interface
	ns         string
}

func (c *moduleClient) List(opts metav1.ListOptions) ([]v1alpha1.Module, error) {
	result := v1alpha1.ModuleList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("modules").
		Do(context.Background()).
		Into(&result)

	return result.Items, err
}

func (c *moduleClient) Get(name string) (*v1alpha1.Module, error) {
	result := v1alpha1.Module{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("modules").
		Name(name).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *moduleClient) Create(project *v1alpha1.Module) (*v1alpha1.Module, error) {
	result := v1alpha1.Module{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("modules").
		Body(project).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

// Update takes the representation of a service and updates it. Returns the server's representation of the service, and an error, if there is any.
func (c *moduleClient) Update(module *v1alpha1.Module) (project *v1alpha1.Module, err error) {
	result := &v1alpha1.Module{}
	err = c.restClient.Put().
		Namespace(c.ns).
		Resource("modules").
		Name(module.Name).
		Body(module).
		Do(context.TODO()).
		Into(result)
	return
}

func (c *moduleClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.restClient.Get().
		Namespace(c.ns).
		Resource("modules").
		Timeout(timeout).
		Watch(context.Background())
}

func (c *moduleClient) Delete(name string) error {
	return c.restClient.
		Delete().
		Namespace(c.ns).
		Resource("modules").
		Name(name).
		Do(context.Background()).
		Error()
}