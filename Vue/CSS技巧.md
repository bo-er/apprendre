1. 使用悬浮层并且支持滑动

   ```css
   .devicePicker {
     position: fixed;
     z-index: 1000;
     width: 100vw;
     height: 100vh;
     left: 0;
     top: 0;
     background-color: #2f2f2fe0;
     align-items: center;
     justify-content: center;
     display: flex;
     
     
     top: 0;
     bottom: 0;
     overflow-y: scroll;
   }
   ```

   关键点是:

   ```
   top: 0;
   bottom: 0;
   overflow-y: scroll;
   ```

   