1. struct转map[string]interface
    ```
    type Data struct {
	    Name string
	    Age int
    }

    func Struct2Map(data Data) {
	    m := make(map[string]interface{})
	    t := reflect.TypeOf(data)
	    v := reflect.ValueOf(data)
	
	    for i := 0; i < t.NumField(); i++ {
		    if !v.Field(i).IsZero() {
			    m[t.Field(i).Name] = v.Field(i).Interface()
    		}
	    }
    }
    ```

2. 在一个给定的有序数组中找两个和为某个定值的数
    ```
    func LookUp(array []int32, targetNum int32) (num1, num2 int32) {
    	left := 0
    	right := len(array) - 1
    	for i := 0; i < len(array); i++ {
    		if array[left]+array[right] > targetNum {
    			right--
    		} else if array[left]+array[right] < targetNum {
    			left++
    		} else {
    			num1, num2 = array[left], array[right]
    			return
    		}
    	}
    	return
    }
    ```

3. 判断两个字符串是否相等
    ```
    func IsStrEqual(s1, s2 string) bool {
    	if len(s1) != len(s2) {
    		return false
    	}
    	for i := range s1 {
    		if s1[i] != s2[i] {
    			return false
    		}
    	}
    	return true
    }
    ```

4. 翻转slice
    ```
    func ReverseSlice(s []interface{}) {
    	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    		s[i], s[j] = s[j], s[i]
    	}
    }
    ```

5. 判断slice中是否存在某个item
    ```
    func IsExistItem(value interface{}, array interface{}) bool {
        switch reflect.TypeOf(array).Kind() {
            case reflect.Slice:
            s := reflect.ValueOf(array)
            for i := 0; i < s.Len(); i++ {
                if reflect.DeepEqual(value, s.Index(i).Interface()) {
                    return true
                }
            }
        }
        return false
    }
    ```

6. 文件拷贝
    ```
    func CopyFile(destName, sourceName string) (int64, error) {
        src, err := os.Open(sourceName)
        if err != nil {
            panic(err)
        }
        defer src.Close()
        
        dest, err := os.Create(destName)
        if err != nil {
            panic(err)
        }
        defer dest.Close()
        return io.Copy(dest, src)
    }
    ```

7. 判断两个map是否中的元素是否一样
    ```
    func IsMapEqual(m1, m2 map[int]int) bool {
    	if len(m1) != len(m2) {
    		return false
    	}
    	for k, m1V := range m1 {
    		if m2V, ok := m2[k]; !ok || m2V != m1V {
    			return false
    		}
    	}
    	return true
    }
    ```

8. 判断字符串切片x是否包含y中的每个元素
    ```
    func IsContainAll(x, y []string) bool {
        for len(y) <= len(x) {
            if len(y) == 0 {
                return true
            }
            if y[0] == x[0] {
                y = y[1:]
            }
            x = x[1:]
        }
    }
    ```

9. 工厂模式
    ```
    import "fmt"
    
    type options struct {
        a int64
        b string
        c map[int]string
    }
    
    func NewOption(opt ...ServerOption) *options {
        r := new(options)
        for _, o := range opt {
            o(r)
        }
        return r
    }
    
    type ServerOption func(*options)
    
    func WriteA(s int64) ServerOption {
        return func(o *options) {
            o.a = s
        }
    }
    
    func WriteB(s string) ServerOption {
        return func(o *options) {
            o.b = s
        }
    }
    
    func WriteC(s map[int]string) ServerOption {
        return func(o *options) {
            o.c = s
        }
    }
    
    func main() {
        opt1 := WriteA(int64(1))
        opt2 := WriteB("test")
        opt3 := WriteC(make(map[int]string, 0))
        op := NewOption(opt1, opt2, opt3)
        fmt.Println(op.a, op.b, op.c)
    }
    ```

10. 元素去重
    ```
    //
    func RemoveRepByLoop(slc []int) []int {
        result := []int{}
        for i := range slc {
            flag := true
            for j := range result {
                if slc[i] == result[j] {
                    flag = false
                    break
                }
            }
            if flag {
                result = append(result, slc[i])
            }
        }
        return result
    }
    
    //通过map转换
    func RemoveRepByMap(slc []int) []int {
    	result := make([]int, 0)
    	tempMap := make(map[int]byte, 0)
    	for _, e := range slc {
    		l := len(tempMap)
    		tempMap[e] = 0
    		if len(tempMap) != l {
    			result = append(result, e)
    		}
    	}
    	return result
    }
    ```
