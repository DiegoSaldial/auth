<template>
  <div class="q-pa-md">
    <Atras />

    <q-form @submit="onSubmit" class="q-mt-lg">
      <div class="row q-gutter-sm justify-center">

        <div class="col-xs-12 col-sm-6">
          <p>Confirmar la eliminacion del conductor</p>
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
import VehiculosServices from '../../../services/taxis/vehiculos';
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
    const vehiculosServices = new VehiculosServices();
    const iduser = ref()
    const idvehi = ref()

    function onSubmit() {
      const data:CreateConductorVehiculos = {
          usuario_id: ''+iduser.value,
          vehiculo_id: idvehi.value
        }
        registrar(data);
    }

    async function registrar(input: CreateConductorVehiculos) {
      const res = await vehiculosServices.quitarVehiculo(input).then((e) => e).catch((e) => e)
      console.log(res);
      if (res.quitarVehiculo) router.back();
    }

    onMounted(()=>{
      iduser.value = route.query.iduser;
      idvehi.value = route.query.idvehi;
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
