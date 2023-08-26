
<template>
  <div class="q-pa-md">
    <Atras />

    <div class="row justify-center">
      <div class="col-xs-12 col-md-12 q-pt-md">
        <q-input v-model.trim="nombre" readonly outlined label="Nombre del rol:" lazy-rules dense
          :rules="[(v) => (v && v.length > 0) || 'Dato requerido.']" />
      </div>
    </div>

    <q-table title="Visualizacion de los permisos de este rol" :loading="loading" :rows="permisos" :columns="columns"
      row-key="name" hide-pagination :rows-per-page-options="[0]" :filter="filter">

      <template v-slot:body-cell-metodo="props">
        <q-td :props="props">
          {{ store.formatCamelCaseString(props.row.metodo) }}
        </q-td>
      </template>
      <template v-slot:body-cell-check="props">
        <q-td :props="props">
          <q-checkbox dense disable v-model="props.row.check" label="" color="green" />
        </q-td>
      </template>
    </q-table>

  </div>
</template>

<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script lang="ts">
import { ref, onMounted } from 'vue';
import PermisosService from '../../services/permisosService';
import { pageStore } from '../../stores/pageStore';
import { PermisosResponseCustom } from '../../types/permisos';
import { useRoute } from 'vue-router';
import Atras from '../utils/header-back.vue'
const columns = [
  { name: 'rol_bits', label: 'Rol Bits', field: 'rol_bits', sortable: true },
  { name: 'metodo', label: 'Permiso', field: 'metodo', sortable: true },
  { name: 'descripcion', label: 'Descripcion', field: 'descripcion', sortable: true },
  { name: 'check', label: 'Acceso', field: 'check', sortable: true },
];

export default {
  name: 'IndexRoles',
  components: { Atras },

  setup() {
    const store = pageStore();
    const permisos = ref<PermisosResponseCustom[]>([]);
    const loading = ref(false);
    const route = useRoute();
    const permisosService = new PermisosService();
    const id = ref()
    const input = ref()
    const nombre = ref('')

    onMounted(async () => {
      loading.value = true;
      id.value = route.query.idrol;
      nombre.value = '' + route.query.n;
      if (id.value) {
        const res = await permisosService.permisos().then((e) => e).catch((e) => e);
        permisos.value = res.permisos;
        const resr = await permisosService.permisosByRol('' + id.value).then((e) => e).catch((e) => e);
        const asignados = resr.permisosByRol;
        permisos.value.forEach(r => {
          const ex = asignados.find((a: any) => a.metodo == r.metodo);
          if (ex) r.check = true;
          else r.check = false;
        })
        // console.log(resr);
        loading.value = false
      }

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
    };
  },
};
</script>
