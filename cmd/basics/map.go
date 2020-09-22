package basics

/**
 * map底层结构实际上就是hash表 map的key可以是任意基础类型和简单的
 */

func gMap(m map[string]string) {
	for key := range m {
		m[key] = "kimo"
	}
}
