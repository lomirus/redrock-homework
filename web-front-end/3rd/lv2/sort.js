let arr = [{ id: 10, name: 'kc' }, { id: 8, name: 'hy' }, { id: 15, name: 'pipi' }, { id: 2, name: 'mama' }]

arr.sort(a, b => a.id - b.id)//[{id:2,name:'mama'},{id:8,name:'hy'},{id:10,name:'kc'},{id:15,name:'pipi'}]