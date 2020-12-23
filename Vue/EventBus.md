The event bus / publish-subscribe pattern, despite the bad press it  sometimes gets, is still an excellent way of getting unrelated sections  of your application to talk to each other. But wait! Before you go waste a few more precious KBs on another library, why not try Vue’s powerful  built-in event bus?

As it turns out, the [event system used in Vue components](https://www.digitalocean.com/community/tutorials/vuejs-events) is just as happy being used on its own.



## Initializing

The first thing you’ll need to do is create the event bus and export  it somewhere so other modules and components can use it. Listen closely. This part might be tricky.

event-bus.js

```javascript
import Vue from 'vue';
export const EventBus = new Vue();
```

 

What do you know? Turns out it wasn’t tricky at all!

All you need to do is import the Vue library and export an instance of it. (In this case, I’ve called it EventBus.) What you’re essentially getting is a component that’s entirely  decoupled from the DOM or the rest of your app. All that exists on it  are its instance methods, so it’s pretty lightweight.



## Using the Event Bus

Now that you’ve created the event bus, all you need to do to use it  is import it in your components and call the same methods that you would use if you were passing messages between parent and child components.

### Sending Events

Say you have a really excited component that feels the need to notify your entire app of how many times it has been clicked whenever someone  clicks on it. Here’s how you would go about implementing that using EventBus.emit(channel: string, payload1: any, …).

-   I’m using a [single-file-component](https://vuejs.org/v2/guide/single-file-components.html) here, but you can use whatever method of creating components you’d like.

PleaseClickMe.vue

```html
<template>
  <div class="pleeease-click-me" @click="emitGlobalClickEvent()"></div>
</template>

<script>
// Import the EventBus we just created.
import { EventBus } from './event-bus.js';

export default {
  data() {
    return {
      clickCount: 0
    }
  },

  methods: {
    emitGlobalClickEvent() {
      this.clickCount++;
      // Send the event on a channel (i-got-clicked) with a payload (the click count.)
      EventBus.$emit('i-got-clicked', this.clickCount);
    }
  }
}
</script>
```

 

### Receiving Events

Now, any other part of your app kind enough to give PleaseClickMe.vue the attention it so desperately craves can import EventBus and listen on the i-got-clicked channel using EventBus.$on(channel: string, callback(payload1,…)).

the-kindly-script.js

```javascript
// Import the EventBus.
import { EventBus } from './event-bus.js';

// Listen for the i-got-clicked event and its payload.
EventBus.$on('i-got-clicked', clickCount => {
  console.log(`Oh, that's nice. It's gotten ${clickCount} clicks! :)`)
});
```

 

-   If you’d only like to listen for the first emission of an event, you can use EventBus.$once(channel: string, callback(payload1,…)).

### Removing Event Listeners

Once a part of your app gets tired of hearing the amount of times PleaseClickMe.vue has been clicked, they can unregister their handler from that channel like so.

```javascript
// Import the EventBus we just created.
import { EventBus } from './event-bus.js';

// The event handler function.
const clickHandler = function(clickCount) {
  console.log(`Oh, that's nice. It's gotten ${clickCount} clicks! :)`)
}

// Listen to the event.
EventBus.$on('i-got-clicked', clickHandler);

// Stop listening.
EventBus.$off('i-got-clicked', clickHandler);
```

 

-   You could also remove **all** listeners for a particular event using EventBus.$off(‘i-got-clicked’) with no callback argument.
-   If you really need to remove every single listener from EventBus, regardless of channel, you can call EventBus.$off() with no arguments at all.