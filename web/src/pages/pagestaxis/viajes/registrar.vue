<template>
  <div class="q-pa-md">
    <Atras />

    <q-form @submit="onSubmit" class="q-mt-lg">
      <div class="row q-gutter-sm justify-center">

        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.number="input.origen_lat" step="0.01" type="number" outlined label="origen_lat:" lazy-rules dense :rules="[
            (v) => (v && v != 0) || 'Dato requerido.',
          ]" />
        </div>
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.number="input.origen_lon" step="0.01" type="number" outlined label="origen_lon:" lazy-rules dense :rules="[
            (v) => (v && v != 0) || 'Dato requerido.',
          ]" />
        </div>
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.number="input.destino_lat" step="0.01" type="number" outlined label="destino_lat:" lazy-rules dense :rules="[
            (v) => (v !== null && v >= 0) || 'Dato requerido.',
          ]" />
        </div>
        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-input v-model.number="input.destino_lon" step="0.01" type="number" outlined label="destino_lon:" lazy-rules dense :rules="[
            (v) => (v !== null && v >= 0) || 'Dato requerido.',
          ]" />
        </div>

        <div class="col-xs-12 col-sm-6 col-md-3">
          <q-select v-model="usuario" :options="usuarios" option-label="nombres" option-value="id" outlined label="Seleccione cliente:" lazy-rules dense :rules="[
            (v) => (v && v.id > 0) || 'Dato requerido.',]">
            <template v-slot:option="scope">
              <q-item v-bind="scope.itemProps">
                <q-item-section avatar>
                  <q-icon name="bookmark" color="green" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ scope.opt.id + ' - ' + scope.opt.nombres + ' ' + scope.opt.apellidos }} [{{ scope.opt.username }}] </q-item-label>
                  <q-item-label caption v-for="(r,key) in scope.opt.roles" :key="key">+ {{ r.nombre }}</q-item-label>
                </q-item-section>
              </q-item>
            </template>
          </q-select>
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

        <div class="col-xs-12 col-sm-9">
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

<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script lang="ts">
import { ref, onMounted } from 'vue';
import {  useRouter } from 'vue-router';
import Atras from '../../utils/header-back.vue';
import CategoriaVServices from '../../../services/taxis/categoria_vehiculos';
import ViajesService from '../../../services/taxis/viajes';
import UsuariosService from '../../../services/usuariosService';
import { pageStore } from '../../../stores/pageStore';
import { CategoriaVehiculosResponse } from '../../../types/taxis/categoria_vehiculos';
import { CrearInput, CreateViajes } from '../../../types/taxis/viajes';

export default {
  name: 'IndexRoles',
  components: { Atras },

  setup() {
    const store = pageStore();
    const loading = ref(false);
    const router = useRouter();
    const categoriaVServices = new CategoriaVServices();
    const usuariosService = new UsuariosService();
    const viajesService = new ViajesService();
    const input = ref<CreateViajes>(CrearInput);
    const categorias = ref<CategoriaVehiculosResponse[]>([])
    const usuarios = ref([]);
    const usuario = ref();
    const id = ref()

    function onSubmit() {
      console.log('input', input.value);
      if(input.value) {
        const data:any = Object.assign({},input.value)
        const inpt:CreateViajes = Object.assign({},input.value)
        inpt.categoria_id = data.categoria_id.id;
        inpt.pasajero_id = usuario.value.id;
        registrar(inpt);
      }
    }

    async function registrar(input: CreateViajes) {
      const res = await viajesService.createViaje(input).then((e) => e).catch((e) => e)
      console.log(res);
      if (res.createViaje) router.back();
    }

    onMounted(async () => {
      loading.value = true;
      const cvs = await categoriaVServices.categoria_vehiculos().then((e) => e).catch((e) => e);
      categorias.value = cvs.categoria_vehiculos;

      const resr = await usuariosService.usuarios().then((e) => e).catch((e) => e);
      usuarios.value = resr.usuarios;
      loading.value = false
    });
    return {
      store,
      loading,
      filter: ref(''),
      id,
      input,
      categorias,
      usuarios,
      usuario,
      onSubmit,
    };
  },
};
</script>
