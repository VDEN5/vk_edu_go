package uniq

import (
	"fmt"
	"strings"
)

/*параметры вызова*/
type Options struct {
	c  bool
	d  bool
	u  bool
	f  int
	ch int
	i  bool
}

func Init(options *Options) {
	options.c, options.d, options.u, options.f, options.ch, options.i = false, false, false, 0, 0, false
}

/*преобразования строк, помогающие работать с утилитой*/
func normal(str string, options Options) string {
	if options.i {
		str = strings.ToLower(str)
	}
	if options.f > 0 {
		if words := strings.Split(str, " "); options.f >= len(words) {
			str = ""
		} else {
			str = strings.Join(words[options.f:], " ")
		}
	}
	if options.ch > 0 {
		if options.ch >= len(str) {
			str = ""
		} else {
			str = str[options.ch:]
		}
	}
	return str
}

/*получение преобразования на основе входных строк и параметров*/
func getString(strs []string, str string, options Options, count int) []string {
	if (options.d && count > 1) || (options.u && count == 1) || (!options.c && !options.d && !options.u) {
		strs = append(strs, str)
	}
	if options.c {
		strs = append(strs, fmt.Sprintf("%d %s", count, str))
	}
	return strs
}

func Uniq(strs []string, options Options) []string {
	if len(strs) == 0 {
		return make([]string, 0)
	}
	result, prev, count := make([]string, 0), strs[0], 1
	for i := 1; i < len(strs); i++ {
		prevCut, curCut := normal(prev, options), normal(strs[i], options)
		if prevCut == curCut {
			count++
			continue
		}
		result, count, prev = getString(result, prev, options, count), 1, strs[i]
	}
	return getString(result, prev, options, count)
}
