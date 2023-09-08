<template>
  <div class="q-pa-md">
    <Atras />

    <q-form @submit="onSubmit" class="q-mt-lg">
      <div class="row justify-center">
        <div class="col-xs-12 col-md-6">
          <q-input v-model.trim="input.nombre" outlined label="Nombre:" lazy-rules dense :rules="[
            (v) => (v && v.length > 0) || 'Dato requerido.',
            (v: any) => (v && v.length < 30) || '30 max permitido.',
          ]" />
        </div>
      </div>

      <div class="row justify-center">
        <q-btn outline color="green" :loading="store.is_loading_page" :label="id ? 'Guardar Cambios' : 'Registrar'"
          size="small" icon="check" class="q-mt-xl" type="submit" />
      </div>
    </q-form>


  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Atras from '../../utils/header-back.vue';
import CaracteristicaVServices from '../../../services/taxis/caracteristicas_vehiculo';
import {CreateCaractericaVehiculos,CrearInput} from '../../../types/taxis/caracteristicas_vehiculo';
import { pageStore } from '../../../stores/pageStore';

export default {
  name: 'IndexRoles',
  components: { Atras },

  setup() {
    const store = pageStore();
    const loading = ref(false);
    const route = useRoute();
    const router = useRouter();
    const caracteristicaVServices = new CaracteristicaVServices();
    const input = ref<CreateCaractericaVehiculos>(CrearInput);
    const id = ref()

    function onSubmit() {
      console.log('input', input.value);
      if(input.value) {
        const data:CreateCaractericaVehiculos = {
          nombre: input.value.nombre,
        }
        if(id.value) data.id = id.value;
        registrar(data);
      }
    }

    async function registrar(input: CreateCaractericaVehiculos) {
      const res = await caracteristicaVServices.createUpdateCaracteristica(input).then((e) => e).catch((e) => e)
      console.log(res);
      if (res.createUpdateCaracteristica) router.back();
    }

    onMounted(async () => {
      loading.value = true;
      id.value = route.query.id;

      if (id.value) {
        const resr = await caracteristicaVServices.caracteristica('' + id.value).then((e) => e).catch((e) => e);
        input.value = resr.caracteristica;
      }
      loading.value = false

    });
    return {
      store,
      loading,
      filter: ref(''),
      id,
      input,
      onSubmit,
    };
  },
};
</script>
