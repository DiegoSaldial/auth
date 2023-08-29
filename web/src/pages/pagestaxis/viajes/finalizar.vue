<template>
  <div class="q-pa-md">
    <Atras />

    <q-form @submit="onSubmit" class="q-mt-lg">
      <div class="row q-gutter-sm justify-center">

        <div class="col-xs-12 col-sm-6">
          <p>Confirmar la finazalizacion del viaje? esto quiere decir que el cliente esta cancelando.</p>
        </div>
      </div>

      <div class="row justify-center">
        <q-btn outline color="green" :loading="store.is_loading_page" label="confirmar"
          size="small" icon="check" class="q-mt-xl" type="submit" />
      </div>
    </q-form>


  </div>
</template>

<script lang="ts">
import { ref,onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Atras from '../../utils/header-back.vue';
import ViajesServices from '../../../services/taxis/viajes';
import { pageStore } from '../../../stores/pageStore';
import { CreateConductorVehiculos } from '../../../types/taxis/vehiculos';

export default {
  name: 'IndexRoles',
  components: { Atras },

  setup() {
    const store = pageStore();
    const loading = ref(false);
    const route = useRoute();
    const router = useRouter();
    const viajesServices = new ViajesServices();
    const idvehi = ref();

    function onSubmit() {
      registrar();
    }

    async function registrar( ) {
      const res = await viajesServices.finalizarViaje(idvehi.value).then((e) => e).catch((e) => e)
      console.log(res);
      if (res.finalizarViaje) router.back();
    }

    onMounted(()=>{
      idvehi.value = route.query.id;
    })

    return {
      store,
      loading,
      filter: ref(''),
      onSubmit,
    };
  },
};
</script>
