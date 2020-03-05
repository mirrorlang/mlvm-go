package mirror

type Atom struct {
	Type             string //null,bool,int,string,pointer,operator
	V_bool           bool
	Name             string
	V_int            int
	V_string         string
	Point_x, Point_y int
	Operator         string
}
