package microconfig

import (
	"io/ioutil"
)

func ParseFile(filename string) (map[string]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	conf := Parse(content)
	return conf, nil
}

func Parse(b []byte) map[string]string {
	m := len(b)
	if m == 0 {
		return nil
	}
	var c byte
	var key, value string
	conf := make(map[string]string)

	k := 0
	for {
		// skip space
		for {
			c = b[k]
			if c > ' ' {
				break
			}
			k++
			if k >= m {
				goto done
			}
		}
		for c == '#' {
			// skip comment
			for {
				c = b[k]
				if c == '\n' {
					break
				}
				k++
				if k >= m {
					goto done
				}
			}
			// skip space
			for {
				c = b[k]
				if c > ' ' {
					break
				}
				k++
				if k >= m {
					goto done
				}
			}
		}
		// parse name
		t := k
		for {
			c = b[k]
			if c <= ' ' || c == '=' {
				break
			}
			k++
			if k >= m {
				goto done
			}
		}
		key = string(b[t:k])
		// skip space
		for {
			c = b[k]
			if (c > ' ' && c != '=') || c == '\n' {
				break
			}
			k++
			if k >= m {
				goto done
			}
		}
		if c != '\n' {
			// parse value
			t := k
			// find end of line
			for {
				c = b[k]
				if c == '\n' {
					break
				}
				k++
				if k >= m {
					break
				}
			}
			// trim right space
			f := k - 1
			for {
				c = b[f]
				if c > ' ' {
					break
				}
				f--
			}
			value = string(b[t : f+1])
			conf[key] = value
			if k >= m {
				goto done
			}
		}
	}
done:
	return conf
}
