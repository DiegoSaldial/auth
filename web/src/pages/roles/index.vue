<template>
  <div class="q-pa-md">
    <q-table title="Roles registrados en el sistema" :loading="loading" :rows="roles" :columns="columns" row-key="name"
      hide-pagination :rows-per-page-options="[0]" :filter="filter">

      <template v-slot:top-right>
        <q-input outlined dense debounce="300" v-model="filter" placeholder="buscar ..." class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>

        <q-btn-group square flat outline>
          <q-btn outline color="green" label="Agregar" size="small" icon="add" to="/adm/roles/registrar" />
        </q-btn-group>
      </template>



      <template v-slot:body-cell-estado="props">
        <q-td :props="props">
          {{ props.row.estado ? 'Activo' : 'Suspendido' }}
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
                <q-item clickable v-ripple :to="'/adm/roles/permisos?idrol=' + props.row.id + '&n=' + props.row.nombre">
                  <q-item-section avatar>
                    <q-icon color="primary" name="visibility" right />
                  </q-item-section>
                  <q-item-section>Permisos</q-item-section>
                </q-item>

                <q-item clickable v-ripple :to="'/adm/roles/registrar?idrol=' + props.row.id + '&n=' + props.row.nombre">
                  <q-item-section avatar>
                    <q-icon color="primary" name="edit" right />
                  </q-item-section>
                  <q-item-section>Editar</q-item-section>
                </q-item>

                <q-item clickable v-ripple :to="'/adm/roles/eliminar?idrol=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="red" name="delete" right />
                  </q-item-section>
                  <q-item-section>Eliminar</q-item-section>
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
import RolesService from '../../services/rolesService';
import { pageStore } from '../../stores/pageStore';
import { RolesResponse } from '../../types/roles';
const columns = [
  { name: 'id', label: 'ID', field: 'id', sortable: true },
  { name: 'nombre', label: 'nombre', field: 'nombre', sortable: true },
  { name: 'bit', label: 'bit', field: 'bit', sortable: true },
  { name: 'permisos', label: 'Permisos', field: 'permisos', sortable: true },
  { name: 'usuarios', label: 'Usuarios', field: 'usuarios', sortable: true },
  { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexRoles',
  setup() {
    const store = pageStore();
    const roles = ref<RolesResponse[]>([]);
    const loading = ref(false);
    const roleservice = new RolesService();

    onMounted(async () => {
      loading.value = true;
      const res = await roleservice.roles(false).then((e) => e).catch((e) => e);
      roles.value = res.roles;
      // console.log(res);
      loading.value = false
    });
    return {
      store,
      columns,
      roles,
      loading,
      filter: ref(''),
    };
  },
};
</script>
