<template>
  <q-layout view="hHh Lpr lff" container style="height: 100vh" class="shadow-0">
    <q-header>
      <q-toolbar>
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />

        <q-toolbar-title>
          {{ store.getsite.name }}
        </q-toolbar-title>

        <q-btn flat :label="datalogin.username" icon="group">
          <q-menu ref="menupefil" class="q-pa-md">
            <div class="text-h6 q-mb-md">
              <!-- <p class="q-ma-none">Menu del usuario</p> -->
              <p class="q-ma-none"> {{ datalogin.nombres }} {{ datalogin.apellidos }} </p>
              <p class="q-ma-none text-green"> <small>{{ datalogin.rol }}</small> </p>

            </div>
            <div class="row no-wrap q-mb-md">
              <div class="column">
                <q-toggle v-model="$q.dark.isActive" label="Cambiar tema oscuro" color="green" />
              </div>

              <div class="column items-center">
                <q-avatar size="84px">
                  <img src="https://www.10duke.com/wp-content/uploads/2022/10/user-access-management.png" />
                </q-avatar>
              </div>
            </div>
            <q-btn color="primary" icon-right="logout" label="Salir" square size="sm" class="full-width"
              @click="logout()" />
          </q-menu>
        </q-btn>

        <q-inner-loading :showing="store.is_loading_page" color="primary" style="background: transparent">
          <q-spinner-dots size="40px" color="yellow" />
        </q-inner-loading>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered :width="250" :breakpoint="700" class="bg-drawer">
      <q-scroll-area class="fit">

        <q-list>
          <q-item-label header>
            Menu de opciones:
          </q-item-label>

          <EssentialLink v-for="link in essentialLinks" :key="link.title" v-bind="link" />
        </q-list>
      </q-scroll-area>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import EssentialLink from 'components/EssentialLink.vue';
import { pageStore } from '../stores/pageStore';
import { useRouter } from 'vue-router';

const linksList = [
  {
    title: 'Usuarios',
    caption: 'Controlar de acceso al sistema',
    icon: 'group',
    link: '/adm/usuarios',
  },
  {
    title: 'roles',
    caption: 'Roles y permisos del sistema',
    icon: 'groups',
    link: '/adm/roles',
  },

];

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink
  },

  setup() {
    const leftDrawerOpen = ref(false);
    const store = pageStore();
    const router = useRouter();
    const datalogin: any = ref({});

    function logout() {
      store.logout();
      router.push('/');
    }

    onMounted(() => {
      datalogin.value = store.getDataLogin();
    })

    return {
      store,
      essentialLinks: linksList,
      leftDrawerOpen,
      datalogin,
      logout,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value
      }
    }
  }
});
</script>
