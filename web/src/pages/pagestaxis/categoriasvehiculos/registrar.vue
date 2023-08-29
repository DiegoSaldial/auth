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
        <div class="col-xs-12 col-md-6">
          <q-input v-model.trim="input.descripcion" outlined label="descripcion:" lazy-rules dense type="textarea" :rules="[
            (v) => (v && v.length > 0) || 'Dato requerido.',
            (v: any) => (v && v.length < 250) || '250 max permitido.',
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
import CategoriaVServices from '../../../services/taxis/categoria_vehiculos';
import {CreateCategoriaVehiculos,CrearInput} from '../../../types/taxis/categoria_vehiculos';
import { pageStore } from '../../../stores/pageStore';

export default {
  name: 'IndexRoles',
  components: { Atras },

  setup() {
    const store = pageStore();
    const loading = ref(false);
    const route = useRoute();
    const router = useRouter();
    const categoriaVServices = new CategoriaVServices();
    const input = ref<CreateCategoriaVehiculos>(CrearInput);
    const id = ref()

    function onSubmit() {
      console.log('input', input.value);
      if(input.value) {
        const data:CreateCategoriaVehiculos = {
          nombre: input.value.nombre,
          descripcion: input.value.descripcion,
        }
        if(id.value) data.id = id.value;
        registrar(data);
      }
    }

    async function registrar(input: CreateCategoriaVehiculos) {
      const res = await categoriaVServices.createUpdateCategoriaVehiculo(input).then((e) => e).catch((e) => e)
      console.log(res);
      if (res.createUpdateCategoriaVehiculo) router.back();
    }

    onMounted(async () => {
      loading.value = true;
      id.value = route.query.id;

      if (id.value) {
        const resr = await categoriaVServices.categoria_vehiculo('' + id.value).then((e) => e).catch((e) => e);
        input.value = resr.categoria_vehiculo;
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
