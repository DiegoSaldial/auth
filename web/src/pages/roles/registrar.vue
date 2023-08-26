<template>
  <div class="q-pa-md">
    <Atras />

    <q-form @submit="onSubmit" class="q-gutter-md">
      <div class="row justify-center">
        <div class="col-xs-12 col-md-12 q-pt-md">
          <q-input v-model.trim="nombre" outlined label="Nombre del rol:" lazy-rules dense :rules="[
            (v) => (v && v.length > 0) || 'Dato requerido.',
            (v: any) => (v && v.length < 30) || '30 max permitido.',
          ]" />
        </div>
      </div>

      <q-table :loading="loading" :rows="permisos" :columns="columns" row-key="name" hide-pagination
        :rows-per-page-options="[0]" :filter="filter">

        <template v-slot:body-cell-metodo="props">
          <q-td :props="props">
            {{ store.formatCamelCaseString(props.row.metodo) }}
          </q-td>
        </template>
        <template v-slot:body-cell-check="props">
          <q-td :props="props">
            <q-checkbox dense v-model="props.row.check" label="" color="green" />
          </q-td>
        </template>

        <template v-slot:body-cell-opt="props">
          <q-td :props="props">
            <q-btn color="green-10" square flat icon="more_vert" size="small">
              <q-menu anchor="top right" self="top left">
                <q-list style="min-width: 110px">
                  <q-item clickable v-ripple :to="'/adm/roles/permisos?idrol=' + props.row.id">
                    <q-item-section avatar>
                      <q-icon color="primary" name="visibility" right />
                    </q-item-section>
                    <q-item-section>Permisos</q-item-section>
                  </q-item>
                </q-list>
              </q-menu>
            </q-btn>
          </q-td>
        </template>
      </q-table>

      <div class="row justify-center">
        <q-btn outline color="green" :loading="store.is_loading_page" :label="id ? 'Guardar Cambios' : 'Registrar'"
          size="small" icon="check" class="q-mt-xl" type="submit" />
      </div>
    </q-form>


  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue';
import PermisosService from '../../services/permisosService';
import RolesService from '../../services/rolesService';
import { pageStore } from '../../stores/pageStore';
import { PermisosResponse, PermisosResponseCustom } from '../../types/permisos';
import { ModificarRol } from '../../types/roles';
import { useRoute, useRouter } from 'vue-router';
import Atras from '../utils/header-back.vue'
const columns = [
  { name: 'rol_bits', label: 'Rol Bits', field: 'rol_bits', sortable: true },
  { name: 'metodo', label: 'Permiso', field: 'metodo', sortable: true },
  { name: 'descripcion', label: 'Descripcion', field: 'descripcion', sortable: true },
  { name: 'check', label: 'Acceso', field: 'check', sortable: true },
  // { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexRoles',
  components: { Atras },

  setup() {
    const store = pageStore();
    const permisos = ref<PermisosResponseCustom[]>([]);
    const loading = ref(false);
    const route = useRoute();
    const router = useRouter();
    const permisosService = new PermisosService();
    const rolesService = new RolesService();
    const id = ref()
    const input = ref()
    const nombre = ref('')

    function onSubmit() {
      const asignados: string[] = [];
      permisos.value.forEach(x => {
        if (x.check) asignados.push(x.metodo)
      });
      const input: ModificarRol = {
        id: id.value,
        nombre: nombre.value,
        permisos: asignados
      }
      console.log('input', input);
      if (!id.value) {
        delete input.id;
        registrar(input);
      } else {
        modificar(input);
      }
    }

    async function registrar(input: ModificarRol) {
      const res = await rolesService.createRol(input).then((e) => e).catch((e) => e)
      console.log(res);
      if (res.createRol) router.back();
    }
    async function modificar(input: ModificarRol) {
      const res = await rolesService.modificarRol(input).then((e) => e).catch((e) => e)
      console.log(res);
      if (res.modificarRol) router.back();
    }

    onMounted(async () => {
      loading.value = true;
      id.value = route.query.idrol;
      nombre.value = '';
      let asignados: PermisosResponse[] = [];
      const res = await permisosService.permisos().then((e) => e).catch((e) => e);
      permisos.value = res.permisos ? res.permisos : [];


      if (id.value) {
        nombre.value = '' + route.query.n;
        const resr = await permisosService.permisosByRol('' + id.value).then((e) => e).catch((e) => e);
        asignados = resr.permisosByRol ? resr.permisosByRol : [];
      }

      permisos.value.forEach(r => {
        const ex = asignados.find(a => a.metodo == r.metodo);
        if (ex) r.check = true;
        else r.check = false;
      })
      loading.value = false

    });
    return {
      store,
      columns,
      permisos,
      loading,
      filter: ref(''),
      id,
      input,
      nombre,
      onSubmit,
    };
  },
};
</script>
