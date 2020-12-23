1. 使用Vuex的时候发现

   ```javascript
   const state = {
   
     stashedDevices: [],
   
     collectingDevices: [],
   
     collectingAttributes: {}
   
   }
   
   const mutations = {
   
     stashNewDevice (state, payload) {
       console.log('stashNewDevice')
       if (!state.stashedDevices){
         state.stashedDevices = []
       }
       //如果不加上面的当state.stashedDevices = []
       //那么下面的代码将报错，提示state.stashedDevices is null
       state.stashedDevices.push(payload)
     },
   }
   ```

   这里怀疑在vuex中定义state内层的数据结构不生效

   

