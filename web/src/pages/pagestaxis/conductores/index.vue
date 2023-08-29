<template>
  <div class="q-pa-md">
    <q-table title="Conductores registrados" :loading="loading" :rows="conductores" :columns="columns"
      row-key="name" hide-pagination :rows-per-page-options="[0]" :filter="filter">

      <template v-slot:top-right>
        <q-input outlined dense debounce="300" v-model="filter" placeholder="buscar ..." class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>

        <!-- <q-btn-group square flat outline>
          <q-btn outline color="green" label="Agregar" size="small" icon="add" to="/adm/vehiculos/registrar" />
        </q-btn-group> -->
      </template>

      <template v-slot:body-cell-foto_url="props">
        <q-td :props="props">
          <q-img :src="props.row.foto_url" spinner-color="purple" style="height: 50px; max-width: 50px"/>
        </q-td>
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
                  :to="'/adm/conductores/vehiculos?userid=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="primary" name="local_taxi" right />
                  </q-item-section>
                  <q-item-section>Vehiculos</q-item-section>
                </q-item>
                <!-- <q-item clickable v-ripple
                  :to="'/adm/vehiculos/asignarConductor?id=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="primary" name="psychology" right />
                  </q-item-section>
                  <q-item-section>Asig. Conductor</q-item-section>
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
import PersonasServices from '../../../services/taxis/personas';
import { pageStore } from '../../../stores/pageStore';
import { UsuariosResponse } from '../../../types/taxis/personas';
const columns = [
  { name: 'id', label: 'ID', field: 'id', sortable: true },
  { name: 'foto_url', label: 'foto_url', field: 'foto_url', sortable: false },
  { name: 'nombres', label: 'nombres', field: 'nombres', sortable: true },
  { name: 'apellidos', label: 'apellidos', field: 'apellidos', sortable: true },
  { name: 'username', label: 'username', field: 'username', sortable: true },
  { name: 'telefono', label: 'telefono', field: 'telefono', sortable: true },
  { name: 'documento', label: 'documento', field: 'documento', sortable: true },
  { name: 'fecha_nac', label: 'fecha_nac', field: 'fecha_nac', sortable: true },
  { name: 'domicilio', label: 'domicilio', field: 'domicilio', sortable: true },
  { name: 'correo', label: 'correo', field: 'correo', sortable: true },
  { name: 'registrado', label: 'registrado', field: 'registrado', sortable: true },
  { name: 'estado', label: 'estado', field: 'estado', sortable: true },
  { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexEmpresas',
  setup() {
    const store = pageStore();
    const conductores = ref<UsuariosResponse[]>([]);
    const loading = ref(false);
    const conductoresServices = new PersonasServices();

    onMounted(async () => {
      loading.value = true;
      const res = await conductoresServices.conductores().then((e) => e).catch((e) => e);
      conductores.value = res.conductores;
      loading.value = false
    });
    return {
      store,
      columns,
      conductores,
      loading,
      filter: ref(''),
    };
  },
};
</script>
