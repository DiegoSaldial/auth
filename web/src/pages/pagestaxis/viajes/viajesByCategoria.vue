<template>
  <div class="q-pa-md">
    <Atras />

    <q-table title="Viajes por categoria" :loading="loading" :rows="viajes" :columns="columns"
      row-key="name" hide-pagination :rows-per-page-options="[0]" :filter="filter">

      <template v-slot:top-right>
        <q-select v-model="categoria" :options="categorias" option-label="nombre" option-value="id" outlined label="Seleccione categoria:" lazy-rules class="full-width" dense :rules="[
            (v) => (v && v.id > 0) || 'Dato requerido.',]">
            <template v-slot:option="scope">
              <q-item v-bind="scope.itemProps">
                <q-item-section avatar>
                  <q-icon name="bookmark" color="green" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ scope.opt.id + ' - ' + scope.opt.nombre }} </q-item-label>
                  <q-item-label caption>{{ scope.opt.descripcion }} </q-item-label>
                </q-item-section>
              </q-item>
            </template>
          </q-select>

        <q-btn-group square flat outline class="q-mr-xl">
          <q-btn outline color="green" label="buscar" size="small" icon="add" @click="listar()" />
        </q-btn-group>

        <q-input outlined dense debounce="300" v-model="filter" placeholder="buscar ..." class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
      </template>

      <template v-slot:body-cell-pasajero_username="props">
        <q-td :props="props">
          <router-link :to="'/adm/usuarios/ver?username='+props.row.pasajero_username" class="text-orange">
            {{ props.row.pasajero_username }}
          </router-link>
        </q-td>
      </template>
      <template v-slot:body-cell-conductor_username="props">
        <q-td :props="props">
          <router-link :to="'/adm/usuarios/ver?username='+props.row.conductor_username" class="text-orange">
            {{ props.row.conductor_username }}
          </router-link>
        </q-td>
      </template>

      <template v-slot:body-cell-estado="props">
        <q-td :props="props">
          {{ props.row.estado?'activo':'finalizado' }}
        </q-td>
      </template>

      <template v-slot:body-cell-origen_lat="props">
        <q-td :props="props">
          <p class="q-my-none"> {{ props.row.origen_lat }} </p>
          <p class="q-my-none"> {{ props.row.origen_lon }} </p>
        </q-td>
      </template>

      <template v-slot:body-cell-destino_lat="props">
        <q-td :props="props">
          <p class="q-my-none"> {{ props.row.destino_lat }} </p>
          <p class="q-my-none"> {{ props.row.destino_lon }} </p>
        </q-td>
      </template>

      <template v-slot:body-cell-registrado="props">
        <q-td :props="props">
          {{ store.parseDate(props.row.registrado) }}
        </q-td>
      </template>

      <template v-slot:body-cell-opt="props">
        <q-td :props="props">
          <q-btn color="green-10" square flat icon="more_vert" size="small">
            <q-menu anchor="top right" self="top left">
              <q-list style="min-width: 110px">
                <!-- <q-item clickable v-ripple :to="'/adm/vehiculos/ver?username=' + props.row.username">
                  <q-item-section avatar>
                    <q-icon color="primary" name="visibility" right />
                  </q-item-section>
                  <q-item-section>Mostrar</q-item-section>
                </q-item> -->

                <q-item clickable v-ripple
                  :to="'/adm/viajes/finalizar?id=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="primary" name="edit" right />
                  </q-item-section>
                  <q-item-section>Finalizar</q-item-section>
                </q-item>
              </q-list>
            </q-menu>
          </q-btn>
        </q-td>
      </template>
    </q-table>
  </div>
</template>

<script lang="ts">
import { onMounted, ref } from 'vue';
import CategoriaVehiService from '../../../services/taxis/categoria_vehiculos';
import ViajesServices from '../../../services/taxis/viajes';
import { pageStore } from '../../../stores/pageStore';
import { Viajes } from '../../../types/taxis/viajes';
import Atras from '../../utils/header-back.vue';

const columns = [
  { name: 'id', label: 'ID', field: 'id', sortable: true },
  // { name: 'pasajero_id', label: 'pasajero_id', field: 'pasajero_id', sortable: true },
  { name: 'conductor_username', label: 'conductor', field: 'conductor_username', sortable: true },
  { name: 'estado', label: 'estado', field: 'estado', sortable: true },
  { name: 'descripcion', label: 'descripcion', field: 'descripcion', sortable: true },
  // { name: 'categoria_id', label: 'categoria_id', field: 'categoria_id', sortable: true },
  { name: 'origen_lat', label: 'origen', field: 'origen_lat', sortable: true },
  // { name: 'origen_lon', label: 'origen_lon', field: 'origen_lon', sortable: true },
  { name: 'destino_lat', label: 'destino', field: 'destino_lat', sortable: true },
  // { name: 'destino_lon', label: 'destino_lon', field: 'destino_lon', sortable: true },
  { name: 'registrado', label: 'registrado', field: 'registrado', sortable: true },
  { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexEmpresas',
  components:{Atras},
  setup() {
    const store = pageStore();
    const viajes = ref<Viajes[]>([]);
    const loading = ref(false);
    const viajesServices = new ViajesServices();
    const categoriaVehiService = new CategoriaVehiService();
    const categorias = ref([]);
    const categoria = ref();

    async function listar(){
      if(!categoria.value) return;
      loading.value = true;
      const res = await viajesServices.viajesByCategoria(categoria.value.id).then((e) => e).catch((e) => e);
      viajes.value = res.viajesByCategoria;
      loading.value = false;
    }

    onMounted(async () => {
      loading.value = true;
      const res = await categoriaVehiService.categoria_vehiculos().then((e) => e).catch((e) => e);
      categorias.value = res.categoria_vehiculos;
      loading.value = false

    });
    return {
      store,
      columns,
      viajes,
      loading,
      categoria,
      categorias,
      filter: ref(''),
      listar,
    };
  },
};
</script>
