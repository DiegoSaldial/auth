<template>
  <div class="q-pa-md">
    <q-table title="Vehiculos registrados" :loading="loading" :rows="vehiculos" :columns="columns"
      row-key="name" hide-pagination :rows-per-page-options="[0]" :filter="filter">

      <template v-slot:top-right>
        <q-input outlined dense debounce="300" v-model="filter" placeholder="buscar ..." class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>

        <q-btn-group square flat outline>
          <q-btn outline color="green" label="Agregar" size="small" icon="add" to="/adm/vehiculos/registrar" />
        </q-btn-group>
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
                  :to="'/adm/vehiculos/registrar?id=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="primary" name="edit" right />
                  </q-item-section>
                  <q-item-section>Editar</q-item-section>
                </q-item>
                <q-item clickable v-ripple v-if="!props.row.conductor_id"
                  :to="'/adm/vehiculos/asignarConductor?id=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="primary" name="psychology" right />
                  </q-item-section>
                  <q-item-section>Asig. Conductor</q-item-section>
                </q-item>
                <q-item clickable v-ripple v-if="props.row.conductor_id"
                  :to="'/adm/vehiculos/quitarConductor?idvehi=' + props.row.id+'&iduser='+props.row.conductor_id">
                  <q-item-section avatar>
                    <q-icon color="red" name="psychology" right />
                  </q-item-section>
                  <q-item-section>Quit. Conductor</q-item-section>
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
import { ref, onMounted } from 'vue';
import VehiculosService from '../../../services/taxis/vehiculos';
import { pageStore } from '../../../stores/pageStore';
import { CategoriaVehiculosResponse } from '../../../types/taxis/categoria_vehiculos';
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
  // { name: 'estado_conductor_vehiculo', label: 'estado_conductor_vehiculo', field: 'estado_conductor_vehiculo', sortable: true },
  { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexEmpresas',
  setup() {
    const store = pageStore();
    const vehiculos = ref<CategoriaVehiculosResponse[]>([]);
    const loading = ref(false);
    const vehiculosService = new VehiculosService();

    onMounted(async () => {
      loading.value = true;
      const res = await vehiculosService.vehiculos().then((e) => e).catch((e) => e);
      vehiculos.value = res.vehiculos;
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
