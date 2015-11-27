package BBB_GPIO

/*import (
        "fmt" "os"
)*/

type bbb_gpio struct {
        pin       [][]int  //Pin name ex.P8_13 Pin(9,12) 
        pin_state map[int]pin_data
        state int
}

//pin location store 

type pin_data struct {
        num_pin    int 
        beagle_pin []int
}

//SIGNAL VALUE 

const (
        OUTPUT = 1 
        INPUT  = 2 
        NONE   = 0

        HIGH   = 1 
        LOW    = 0

        set = 1 
        none = 0 
        ERROR = -1
)

//mode set 

var mode = map[int]string{
        1: "out", 2: "in",
}

//pin allocate

var pin_map = map[int][]int{
        8: []int{0, 0, 0, 38, 39, 34, 35, 66, 67, 69, 68, 45, 44, 23,
        26, 47, 46, 27, 65, 22, 63, 62, 37, 36, 33, 32, 61, 86, 88, 87,
        89, 10, 11, 9, 81, 8, 80, 78, 79, 76, 77, 74, 75, 72, 73, 70,
        71},
        9: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 30, 60, 31, 50,
        48, 51, 5, 4, 0, 0, 3, 2, 49, 15, 117, 14, 115, 123, 121, 122,
        120, 0, 0, 0, 0, 0, 0, 0, 0, 20, 7, 0, 0, 0, 0},
}

//default structure 

var gpio bbb_gpio = bbb_gpio{}

func (e *bbb_gpio) Error() string{
	if e.state == ERROR{
		return "BBB_GPIO Error"
	}
        return "BBB_GPIO Exception"
}

func Start()  (*bbb_gpio,error){
	if gpio.state != set && gpio.state != ERROR{
		gpio.pin = make([][]int,10)
                 
                gpio.pin[8] =pin_map[8] 
                gpio.pin[9] = pin_map[9]
                 
                gpio.pin_state = make(map[int]pin_data)
	}else if gpio.state == set{
		return &gpio, &gpio
	}else {
		return nil, &gpio
	}
        return &gpio, nil
}
