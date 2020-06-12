/* eslint-disable no-console */

import { register } from 'register-service-worker';
import { Notify } from 'quasar';

if (process.env.NODE_ENV === 'production') {
  register(`${process.env.BASE_URL}service-worker.js`, {
    ready() {
      console.log(
        'App is being served from cache by a service worker.\n' +
          'For more details, visit https://goo.gl/AFskqB'
      );
    },
    registered() {
      console.log('Service worker has been registered.');
    },
    cached() {
      console.log('Content has been cached for offline use.');
    },
    updatefound() {
      console.log('New content is downloading.');
    },
    updated() {
      console.log('New content is available; please refresh.');
      // window.location.reload();
      // Notify.create({
      //   color: 'negative',
      //   icon: 'cached',
      //   message: 'Updated content is available. Please refresh the page.',
      //   timeout: 0,
      //   multiLine: true,
      //   position: 'top',
      //   actions: [
      //     {
      //       label: 'Refresh',
      //       color: 'yellow',
      //       handler: () => {
      //         window.location.reload(true);
      //       },
      //     },
      //     {
      //       label: 'Dismiss',
      //       color: 'white',
      //       handler: () => {},
      //     },
      //   ],
      // });
    },
    offline() {
      console.log(
        'No internet connection found. App is running in offline mode.'
      );
    },
    error(error) {
      console.error('Error during service worker registration:', error);
    },
  });
}
