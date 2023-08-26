<template>
  <div class="">
    <q-layout view="hHh Lpr lff" container style="height: 100vh" class="shadow-0">
      <q-header class="bg-primary">
        <q-toolbar>
          <q-toolbar-title>
            <q-img :src="store.getsite?.logo" alt="logo" :ratio="store.getsite?.ratio" style="
                height: 30px;
                max-width: 120px;
                margin-right: 0.5em;
                border-radius: 0%;
              " />
            <span style="margin-right: 0.7em" @click="$router.push('/')">
              Acceso al Sistema
            </span>
          </q-toolbar-title>
          <q-toggle v-model="$q.dark.isActive" color="white" />
          <q-inner-loading :showing="store.is_loading_page" color="primary" style="background: transparent">
            <q-spinner-dots size="40px" color="white" />
          </q-inner-loading>
        </q-toolbar>
      </q-header>

      <LoginForm />
    </q-layout>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { pageStore } from '../stores/pageStore';
import LoginForm from '../pages/xauth/login-form.vue';

export default defineComponent({
  name: 'LoginLayout',

  components: {
    LoginForm,
  },

  setup() {
    const store = pageStore();
    return {
      store,
    };
  },
});
</script>
