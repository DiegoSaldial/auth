<template>
  <div class="q-pa-md">
    <Atras />

    <q-form @submit="onSubmit" class="q-mt-lg">
      <div class="row q-gutter-sm justify-center">
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.trim="input.placa" outlined label="Placa:" lazy-rules dense :rules="[
            (v) => (v && v.length > 0) || 'Dato requerido.',
            (v: any) => (v && v.length < 15) || '15 max permitido.',
          ]" />
        </div>
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.number="input.puertas" type="number" outlined label="Puertas:" lazy-rules dense :rules="[
            (v) => (v !== null && v >= 0) || 'Dato requerido.',
          ]" />
        </div>
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.number="input.capacidad" type="number" outlined label="Capacidad:" lazy-rules dense :rules="[
            (v) => (v && v > 0) || 'Dato requerido.',
          ]" />
        </div>
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.trim="input.color" outlined label="Color:" lazy-rules dense :rules="[
            (v) => (v && v.length > 0) || 'Dato requerido.',
            (v: any) => (v && v.length < 30) || '30 max permitido.',
          ]" />
        </div>
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.trim="input.modelo" outlined label="Modelo:" lazy-rules dense :rules="[
            (v) => (v && v.length > 0) || 'Dato requerido.',
            (v: any) => (v && v.length < 30) || '30 max permitido.',
          ]" />
        </div>
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.number="input.anio" type="number" outlined label="Año:" lazy-rules dense :rules="[
            (v) => (v && v > 1900) || 'Dato requerido mayor a 1900.',
          ]" />
        </div>
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-select v-model="input.categoria_id" :options="categorias" option-label="nombre" option-value="id" outlined label="Categoria:" lazy-rules dense :rules="[
            (v) => (v && v.id > 0) || 'Dato requerido.',]">
            <template v-slot:option="scope">
              <q-item v-bind="scope.itemProps">
                <q-item-section avatar>
                  <q-icon name="bookmark" color="green" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ scope.opt.id + ' - ' + scope.opt.nombre }}</q-item-label>
                  <q-item-label caption>{{ scope.opt.descripcion }}</q-item-label>
                </q-item-section>
              </q-item>
            </template>
          </q-select>
        </div>

        <div class="col-xs-12 col-sm-6">
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
import VehiculosServices from '../../../services/taxis/vehiculos';
import { pageStore } from '../../../stores/pageStore';
import { CreateVehiculos, CreateVehiculosInput } from '../../../types/taxis/vehiculos';
import { CategoriaVehiculosResponse } from '../../../types/taxis/categoria_vehiculos';

export default {
  name: 'IndexRoles',
  components: { Atras },

  setup() {
    const store = pageStore();
    const loading = ref(false);
    const route = useRoute();
    const router = useRouter();
    const categoriaVServices = new CategoriaVServices();
    const vehiculosServices = new VehiculosServices();
    const input = ref<CreateVehiculos>(CreateVehiculosInput);
    const categorias = ref<CategoriaVehiculosResponse[]>([])
    const id = ref()

    function onSubmit() {
      console.log('input', input.value);
      if(input.value) {
        const data:CreateVehiculos = Object.assign({},input.value)
        if(id.value) data.id = id.value;
        data.categoria_id = input.value.categoria_id.id;
        registrar(data);
      }
    }

    async function registrar(input: CreateVehiculos) {
      const res = await vehiculosServices.createUpdateVehiculo(input).then((e) => e).catch((e) => e)
      console.log(res);
      if (res.createUpdateVehiculo) router.back();
    }

    onMounted(async () => {
      loading.value = true;
      id.value = route.query.id;
      const cvs = await categoriaVServices.categoria_vehiculos().then((e) => e).catch((e) => e);
      categorias.value = cvs.categoria_vehiculos;

      if (id.value) {
        const resr = await vehiculosServices.vehiculo('' + id.value).then((e) => e).catch((e) => e);
        input.value = resr.vehiculo;
        const cate:any = categorias.value.find(c=>c.id==input.value.categoria_id)
        if(cate) input.value.categoria_id = cate;
      }
      loading.value = false

    });
    return {
      store,
      loading,
      filter: ref(''),
      id,
      input,
      categorias,
      onSubmit,
    };
  },
};
</script>
