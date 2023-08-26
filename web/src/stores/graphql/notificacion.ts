/* eslint-disable @typescript-eslint/no-explicit-any */
import { Notify } from 'quasar';

export function notificar(men: string, caption = '', color = 'primary') {
  Notify.create({
    message: men,
    caption: caption,
    icon: 'error',
    color: color,
    position: 'bottom-right',
    actions: [
      {
        label: 'Ok',
        color: 'white',
        handler: () => {
          /* ... */
        },
      },
    ],
  });
}
