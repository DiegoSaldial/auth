<template>
  <div class="q-pa-md">
    <Atras />

    <q-table title="Vehiculos por caracteristica seleccionada" :loading="loading" :rows="vehiculos" :columns="columns"
      row-key="name" hide-pagination :rows-per-page-options="[0]" :filter="filter">

      <template v-slot:top-right>
        <q-input outlined dense debounce="300" v-model="filter" placeholder="buscar ..." class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
      </template>

      <template v-slot:body-cell-estado="props">
        <q-td :props="props">
          {{ props.row.estado?'activo':'inactivo' }}
        </q-td>
      </template>

      <template v-slot:body-cell-registrado="props">
        <q-td :props="props">
          {{ store.parseDate(props.row.registrado) }}
        </q-td>
      </template>
    </q-table>
  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue';
import VehiculosService from '../../../services/taxis/vehiculos';
import { pageStore } from '../../../stores/pageStore';
import { CategoriaVehiculosResponse } from '../../../types/taxis/categoria_vehiculos';
import { useRoute } from 'vue-router';
import Atras from '../../utils/header-back.vue';

const columns = [
  { name: 'id', label: 'ID', field: 'id', sortable: true },
  { name: 'categoria', label: 'categoria', field: 'categoria', sortable: true },
  { name: 'placa', label: 'placa', field: 'placa', sortable: true },
  { name: 'puertas', label: 'puertas', field: 'puertas', sortable: true },
  { name: 'capacidad', label: 'capacidad', field: 'capacidad', sortable: true },
  { name: 'descripcion', label: 'descripcion', field: 'descripcion', sortable: true },
  { name: 'color', label: 'color', field: 'color', sortable: true },
  { name: 'modelo', label: 'modelo', field: 'modelo', sortable: true },
  { name: 'anio', label: 'anio', field: 'anio', sortable: true },
  { name: 'estado', label: 'estado veh.', field: 'estado', sortable: true },
  { name: 'registrado', label: 'registrado', field: 'registrado', sortable: true },
  { name: 'conductor', label: 'conductor', field: 'conductor', sortable: true },
  // { name: 'estado_conductor_vehiculo', label: 'estado_conductor_vehiculo', field: 'estado_conductor_vehiculo', sortable: true },
  // { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexEmpresas',
  components:{Atras},
  setup() {
    const store = pageStore();
    const vehiculos = ref<CategoriaVehiculosResponse[]>([]);
    const loading = ref(false);
    const route = useRoute();
    const vehiculosService = new VehiculosService();

    onMounted(async () => {
      loading.value = true;
      const cateid = ''+route.query.id;
      const res = await vehiculosService.vehiculosByCaracteristica(cateid).then((e) => e).catch((e) => e);
      vehiculos.value = res.vehiculosByCaracteristica;
      loading.value = false
    });
    return {
      store,
      columns,
      vehiculos,
      loading,
      filter: ref(''),
    };
  },
};
</script>
