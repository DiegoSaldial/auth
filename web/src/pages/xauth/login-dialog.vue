<template>
  <q-dialog v-model="alert" transition-show="slide-up" transition-hide="slide-down" maximized>
    <q-card>
      <q-card-actions align="right">
        <q-btn flat label="Cerrar" icon="close" color="primary" v-close-popup />
      </q-card-actions>

      <LoginForm :redirect="false" v-on:onlogin="onlogin" />
    </q-card>
  </q-dialog>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import LoginForm from './login-form.vue';

export default defineComponent({
  name: 'LoginDialog',
  components: {
    LoginForm,
  },
  setup() {
    const alert = ref(false);

    function opendialog() {
      alert.value = true;
    }

    function onlogin() {
      alert.value = false;
    }
    return {
      alert,
      opendialog,
      onlogin,
    };
  },
});
</script>
