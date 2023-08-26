<template>
  <div class="q-pa-md">
    <xheader :title="''" />

    <q-separator inset class="q-mb-lg" />
    <q-linear-progress rounded stripe size="2px" indeterminate color="green" class="q-mt-sm"
      v-if="store.is_loading_page" />

    <div class="row justify-center">
      <q-btn outline color="red" :loading="store.is_loading_page" label="Proceder" size="small" icon="delete"
        class="q-mt-xl" @click="eliminar()" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { pageStore } from '../../stores/pageStore';
import RolesService from '../../services/rolesService';
import xheader from '../utils/header-back.vue';

export default defineComponent({
  name: 'VerUsuario',
  components: {
    xheader,
  },
  setup() {
    const id = ref();
    const store = pageStore();
    const route = useRoute();
    const router = useRouter();
    const rolesService = new RolesService();

    async function eliminar() {
      const res = await rolesService.deleteRol(id.value).then((e) => e).catch((e) => e)
      if (res.deleteRol) router.back()
    }

    onMounted(async () => {
      id.value = route.query.idrol;
    });

    return {
      store,
      id,
      eliminar
    };
  },
});
</script>
