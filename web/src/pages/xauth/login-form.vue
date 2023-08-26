<template>
  <q-page-container>
    <h5 class="text-center q-mb-xs">
      Ingrese sus credenciales de acceso al sistema
    </h5>

    <div class="row justify-center">
      <div class="col-12 col-sm-4 col-lg-2">
        <Vue3Lottie :animationData="AstronautJSON3" :pauseOnHover="false" height="50%" width="80%" />
      </div>
    </div>
    <q-form ref="formlogin">
      <div class="row justify-center">
        <div class="col-12 col-sm-5 col-lg-3">
          <q-input v-model.trim="data.username" outlined label="Usuario:" lazy-rules counter :disable="loading" :rules="[
            (v) => (v && v.length > 0) || 'Dato requerido.',
            (v) => (v && v.length < 30) || '30 max permitido.',
          ]">
            <template v-slot:prepend> <q-icon name="people" /> </template>
          </q-input>

          <q-input v-model="data.password" outlined counter :type="isPwd ? 'password' : 'text'" label="Contraseña:"
            :disable="loading" :rules="[
              (v) => (v && v.length > 0) || 'Dato requerido.',
              (v) => (v && v.length < 30) || '30 max permitido.',
            ]" v-on:keyup.enter="login()">
            <template v-slot:prepend> <q-icon name="lock" /> </template>

            <template v-slot:append>
              <q-icon :name="isPwd ? 'visibility_off' : 'visibility'" class="cursor-pointer" @click="isPwd = !isPwd" />
            </template>
          </q-input>

          <q-select v-if="roles && roles.length > 0" outlined v-model="rol" :options="roles" :disable="loading"
            option-label="nombre" option-value="id" label="Seleccione el rol:" />

          <q-linear-progress rounded stripe size="3px" indeterminate color="green" class="q-mt-sm" v-if="loading" />
        </div>
      </div>

      <div class="row justify-center">
        <q-btn outline color="green" :disable="loading" :loading="loading" label="Ingresar" icon="login" class="q-mt-xl"
          @click="login()" />
      </div>
    </q-form>
  </q-page-container>
</template>

<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { pageStore } from '../../stores/pageStore';
import { Vue3Lottie } from 'vue3-lottie';
import { LoginData, Instancia, ResponseLogin } from '../../types/login';
import { useRouter } from 'vue-router';
import LoginService from '../../services/loginService';
import AstronautJSON3 from '../../assets/102642-contact.json';

export default defineComponent({
  name: 'LoginForm',
  props: {
    redirect: { type: Boolean, default: true },
  },
  components: {
    Vue3Lottie,
  },
  setup(props, vue) {
    const store = pageStore();
    const loading = ref(false);
    const isPwd = ref(true);
    const router = useRouter();
    const loginService = new LoginService();
    const formlogin = ref();
    const data = ref<LoginData>(Instancia);
    const rol = ref();
    const roles = ref<ResponseLogin[]>();
    const data_login = ref();

    async function login() {
      const v = await formlogin.value.validate().then((r: any) => r);
      if (!v) return;
      data_login.value = null;
      if (rol.value) {
        loginWithRol();
        return
      }
      loading.value = true;
      const res = await loginService.login(data.value).then((e) => e).catch((e) => e)
      // console.log(res, rol.value);
      if (res.login) {
        roles.value = res.login;
      }
      loading.value = false;
    }

    async function loginWithRol() {
      loading.value = true;
      const idrol = rol.value.id;
      const res = await loginService.useRol(data.value, idrol).then((e) => e).catch((e) => e)
      // console.log(res);
      if (res.useRol) {
        roles.value = res.useRol;
        data_login.value = JSON.parse(JSON.stringify(res.useRol));
        data_login.value.password = data.value.password;
        store.setDataLogin(data_login.value);
        if (props.redirect) router.push('/adm');
        vue.emit('onlogin');
      } else {
        loading.value = false;
      }
    }

    onMounted(() => {
      data.value.username = '';
      data.value.password = '';
    });

    return {
      store,
      loading,
      isPwd,
      formlogin,
      data,
      AstronautJSON3,
      rol,
      roles,
      login,
    };
  },
});
</script>
