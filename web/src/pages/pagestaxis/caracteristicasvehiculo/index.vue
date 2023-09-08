<template>
  <div class="q-pa-md">
    <q-table title="Caracteristicas de vehiculos" :loading="loading" :rows="categorias" :columns="columns"
      row-key="name" hide-pagination :rows-per-page-options="[0]" :filter="filter">

      <template v-slot:top-right>
        <q-input outlined dense debounce="300" v-model="filter" placeholder="buscar ..." class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>

        <q-btn-group square flat outline>
          <q-btn outline color="green" label="Agregar" size="small" icon="add" to="/adm/caracteristicas/registrar" />
        </q-btn-group>
      </template>

      <template v-slot:body-cell-vehiculos="props">
        <q-td :props="props">
          <router-link :to="'/adm/caracteristicas/vehiculosByCaracteristica?id='+props.row.id" class="text-orange">
            {{ props.row.vehiculos }}
          </router-link>
        </q-td>
      </template>

      <template v-slot:body-cell-registro="props">
        <q-td :props="props">
          {{ store.parseDate(props.row.registro) }}
        </q-td>
      </template>

      <template v-slot:body-cell-opt="props">
        <q-td :props="props">
          <q-btn color="green-10" square flat icon="more_vert" size="small">
            <q-menu anchor="top right" self="top left">
              <q-list style="min-width: 110px">
                <q-item clickable v-ripple
                  :to="'/adm/caracteristicas/registrar?id=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="primary" name="edit" right />
                  </q-item-section>
                  <q-item-section>Editar</q-item-section>
                </q-item>
                <!-- <q-item clickable v-ripple
                  :to="'/adm/categorias_vehiculos/vehiculosByCategoria?id=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="primary" name="local_taxi" right />
                  </q-item-section>
                  <q-item-section>Vehiculos</q-item-section>
                </q-item> -->
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
import CaracteristicaVehiService from '../../../services/taxis/caracteristicas_vehiculo';
import { pageStore } from '../../../stores/pageStore';
import { CategoriaVehiculosResponse } from '../../../types/taxis/categoria_vehiculos';
const columns = [
  { name: 'id', label: 'ID', field: 'id', sortable: true },
  { name: 'nombre', label: 'Nombre', field: 'nombre', sortable: true },
  { name: 'registro', label: 'Registrado', field: 'registro', sortable: true },
  { name: 'vehiculos', label: 'Vehiculos', field: 'vehiculos', sortable: true },
  { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexEmpresas',
  setup() {
    const store = pageStore();
    const categorias = ref<CategoriaVehiculosResponse[]>([]);
    const loading = ref(false);
    const caracteristicaVehiService = new CaracteristicaVehiService();

    onMounted(async () => {
      loading.value = true;
      const res = await caracteristicaVehiService.caracteristicas().then((e) => e).catch((e) => e);
      categorias.value = res.caracteristicas;
      loading.value = false
    });
    return {
      store,
      columns,
      categorias,
      loading,
      filter: ref(''),
    };
  },
};
</script>
