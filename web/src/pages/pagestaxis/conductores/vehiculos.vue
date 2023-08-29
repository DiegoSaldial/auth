<template>
  <div class="q-pa-md">
    <Atras />
    <q-table title="Vehiculos asignados a este usuario" :loading="loading" :rows="vehiculos" :columns="columns"
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
      <template v-slot:body-cell-estado_conductor_vehiculo="props">
        <q-td :props="props">
          {{ props.row.estado_conductor_vehiculo?'Si':'No' }}
        </q-td>
      </template>
      <template v-slot:body-cell-registrado="props">
        <q-td :props="props">
          {{ store.parseDate(props.row.registrado) }}
        </q-td>
      </template>

      <!-- <template v-slot:body-cell-opt="props">
        <q-td :props="props">
          <q-btn color="green-10" square flat icon="more_vert" size="small">
            <q-menu anchor="top right" self="top left">
              <q-list style="min-width: 110px">
                <q-item clickable v-ripple
                  :to="'/adm/vehiculos/asignarConductor?id=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="primary" name="psychology" right />
                  </q-item-section>
                  <q-item-section>Asig. Conductor</q-item-section>
                </q-item>
              </q-list>
            </q-menu>
          </q-btn>
        </q-td>
      </template> -->
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
  { name: 'categoria', label: 'categoria', field: 'categoria', sortable: true },
  { name: 'estado_conductor_vehiculo', label: 'Conduce actualmente', field: 'estado_conductor_vehiculo', sortable: true },
  { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexEmpresas',
  components:{Atras},
  setup() {
    const store = pageStore();
    const route = useRoute();
    const vehiculos = ref<CategoriaVehiculosResponse[]>([]);
    const vehiculosService = new VehiculosService();
    const loading = ref(false);
    const userid = ref();

    onMounted(async () => {
      loading.value = true;
      userid.value = route.query.userid;
      if(userid.value){
        const res = await vehiculosService.vehiculosByConductor(userid.value).then((e) => e).catch((e) => e);
        vehiculos.value = res.vehiculosByConductor;
      }
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
