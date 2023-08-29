<template>
  <div class="q-pa-md">
    <xheader :title="''" />

    <q-separator inset class="q-mb-lg" />
    <q-linear-progress rounded stripe size="2px" indeterminate color="green" class="q-mt-sm" v-if="store.is_loading_page" />

    <div class="row justify-center q-mt-md">
      <div class="col-12 col-md-6 q-mb-md">
        <xobjeto :objeto="edit" />
      </div>
      <div v-if="edit?.foto_url" class="col-12 col-md-6 q-px-xl">
        <q-img :src="edit?.foto_url" spinner-color="purple" xstyle="height: 50px; max-width: 50px" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import xheader from '../../utils/header-back.vue';
import xobjeto from '../../utils/objeto-view.vue';
import { pageStore } from '../../../stores/pageStore';
import VehiculosServices from '../../../services/taxis/vehiculos';

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
    const edit = ref();
    const vehiculosServices = new VehiculosServices();

    onMounted(async () => {
      id.value = route.query.id;
      if (id.value) {
        const res = await vehiculosServices.vehiculo(id.value).then((e) => e).catch((e) => e);
        edit.value = res.vehiculo;
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
