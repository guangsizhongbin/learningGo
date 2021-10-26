package concise

// new:
// 内建函数 new 本质上说跟其他语言中的同名函数功能一样: new(T)分配了零值填充的T类型的内在空间
// 并且返回其地址, 即一个 *T 类型的值. 用Go语言的术语说，它返回了一个指针，指向新分配的类型T的零值

// make只能创建slice, map, channel, 并且返回一个有初始值(非零)的T类型,而不是*T
