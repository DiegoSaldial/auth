<template>
  <div class="q-pa-md">
    <xheader :title="''" />

    <q-separator inset class="q-mb-lg" />
    <q-linear-progress rounded stripe size="2px" indeterminate color="green" class="q-mt-sm"
      v-if="store.is_loading_page" />

    <div class="row justify-center q-mt-md">
      <div class="col-12 col-md-6">
        <xobjeto :objeto="edit" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import xheader from '../utils/header-back.vue';
import xobjeto from '../utils/objeto-view.vue';
import { UsuarioResponse } from '../../types/usuarios';
import { pageStore } from '../../stores/pageStore';
import UsuariosService from '../../services/usuariosService';

export default defineComponent({
  name: 'VerUsuario',
  components: {
    xheader,
    xobjeto,
  },
  setup() {
    const id = ref();
    const store = pageStore();
    const route = useRoute();
    const edit = ref<UsuarioResponse>();
    const usuarioService = new UsuariosService();

    onMounted(async () => {
      id.value = route.query.username;
      if (id.value) {
        const res = await usuarioService.usuarioByUsername(id.value).then((e) => e).catch((e) => e);
        edit.value = res.usuarioByUsername;
      }
    });

    return {
      store,
      id,
      edit,
    };
  },
});
</script>
