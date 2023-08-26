<template>
  <div class="q-pa-md">
    <q-table title="Usuarios registrados en el sistema" :loading="loading" :rows="usuarios" :columns="columns"
      row-key="name" hide-pagination :rows-per-page-options="[0]" :filter="filter">

      <template v-slot:top-right>
        <q-input outlined dense debounce="300" v-model="filter" placeholder="buscar ..." class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>

        <q-btn-group square flat outline>
          <q-btn outline color="green" label="Agregar" size="small" icon="add" to="/adm/usuarios/registrar" />
        </q-btn-group>
      </template>

      <template v-slot:body-cell-nombres="props">
        <q-td :props="props">
          {{ props.row.nombres }} {{ props.row.apellidos }}
        </q-td>
      </template>
      <template v-slot:body-cell-roles="props">
        <q-td :props="props" style="text-align: left;">
          <p v-for="r in props.row.roles" :key="r.id" class="q-mb-none">+ {{ r.nombre }}</p>
        </q-td>
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
                <q-item clickable v-ripple :to="'/adm/usuarios/ver?username=' + props.row.username">
                  <q-item-section avatar>
                    <q-icon color="primary" name="visibility" right />
                  </q-item-section>
                  <q-item-section>Mostrar</q-item-section>
                </q-item>

                <q-item clickable v-ripple
                  :to="'/adm/usuarios/registrar?iduser=' + props.row.id + '&uname=' + props.row.username">
                  <q-item-section avatar>
                    <q-icon color="primary" name="edit" right />
                  </q-item-section>
                  <q-item-section>Editar</q-item-section>
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
import UsuarioService from '../../services/usuariosService';
import { pageStore } from '../../stores/pageStore';
import { UsuarioLoginResponse } from '../../types/usuarios';
const columns = [
  { name: 'id', label: 'ID', field: 'id', sortable: true },
  { name: 'username', label: 'Username', field: 'username', sortable: true },
  { name: 'nombres', label: 'Usuario', field: 'nombres', sortable: true },
  // { name: 'nombres', label: 'Nombres', field: 'nombres', sortable: true },
  // { name: 'apellidos', label: 'Apellidos', field: 'apellidos', sortable: true },
  { name: 'correo', label: 'Correo', field: 'correo', sortable: true },
  { name: 'roles', label: 'Roles', field: 'roles', sortable: true },
  { name: 'telefono', label: 'Telefono', field: 'telefono', sortable: true },
  { name: 'estado', label: 'Estado', field: 'estado' },
  { name: 'registrado', label: 'Registrado', field: 'registrado', sortable: true },
  { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexEmpresas',
  setup() {
    const store = pageStore();
    const usuarios = ref<UsuarioLoginResponse[]>([]);
    const loading = ref(false);
    const usuarioService = new UsuarioService();

    onMounted(async () => {
      loading.value = true;
      const res = await usuarioService.usuarios().then((e) => e).catch((e) => e);
      usuarios.value = res.usuarios;
      // console.log(res);
      loading.value = false
    });
    return {
      store,
      columns,
      usuarios,
      loading,
      filter: ref(''),
    };
  },
};
</script>
