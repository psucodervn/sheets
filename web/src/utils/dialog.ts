import { Notify } from 'quasar';

type Position =
  | 'top-left'
  | 'top-right'
  | 'bottom-left'
  | 'bottom-right'
  | 'top'
  | 'bottom'
  | 'left'
  | 'right'
  | 'center';

export const showSuccess = (
  message: string,
  caption?: string,
  position: Position = 'bottom'
) => {
  Notify.create({
    message,
    caption,
    position,
    type: 'positive',
    actions: [
      {
        label: 'Dismiss',
        color: 'white',
        handler: () => {
          /* ... */
        },
      },
    ],
  });
};

export const showFailure = (
  message: string,
  caption?: string,
  position: Position = 'bottom'
) => {
  Notify.create({
    message,
    caption,
    position,
    type: 'negative',
    actions: [
      {
        label: 'Dismiss',
        color: 'white',
        handler: () => {
          /* ... */
        },
      },
    ],
  });
};
