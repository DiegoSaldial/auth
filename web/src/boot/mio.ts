import { boot } from 'quasar/wrappers';
import { Cookies } from 'quasar';
import { Dark } from 'quasar';
import { watch } from 'vue';

export default boot(({}) => {
  const value = Cookies.get('quasar-theme-invent');
  Dark.set(value == 'true');

  watch(
    () => Dark.isActive,
    (val) => {
      Dark.set(val);
      Cookies.set('quasar-theme-invent', '' + val);
      // console.log(val ? 'On dark mode' : 'On light mode');
    }
  );
});
