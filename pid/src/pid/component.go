package pid

// An IO component is a component in a closed system.
type IOComponent interface {
  // SetInput sets the input value for the component.
  SetInput(float64)
  // Input gets the input value for the component.
  Input() float64
  // Output gets the output value for the component.
  Output() float64
  // Parameters gets the settable parameters of the component.
  Parameters() []parameter
  // SetParameters sets the parameters of the component.
  SetParameters([]parameter)
}

// A parameter is a configurable parameter for an IOComponent.
type parameter struct {
  // Name is the shortname of the parameter.
  Name string
  // Title is a more human readable name (for UI display)
  Title string
  // Minimum value of the parameter.
  Minimum float64
  // Maximum value of the parameter.
  Maximum float64
  // The desired step size of the parameter.
  Step float64
  // The default value.
  Default float64
  // The units to display.
  Unit string
  // The value (for setting parameters)
  Value float64
}

type Driver interface {
  IOComponent
}

type Load interface {
  IOComponent
}

type Sensor interface {
  IOComponent
}

// SetComponentDefaults sets the parameters of the component to its defaults.
func SetComponentDefaults(c IOComponent) {
  params := c.Parameters()
  for i := 0 ; i < len(params) ; i++ {
    params[i].Value = params[i].Default
  }
  c.SetParameters(params)
}

// A SystemGenerator is registered on init, and is used to build a complete
// System struct when a simulation is requested.
type SystemGenerator interface {
  // Name returns the name of the Generator.
  Name() string
  // GenerateSystem returns a System.
  GenerateSystem() System
}

