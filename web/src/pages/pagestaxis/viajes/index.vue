<template>
  <div class="q-pa-md">
    <q-btn-group square flat outline class="q-mb-xl">
      <q-btn outline color="green" label="Registrar viaje" size="small" icon="add" to="/adm/viajes/registrar" />
    </q-btn-group>

    <div class="q-gutter-y-md">
      <p class="text-center q-ma-none"> <b>Menu de filtro de viajes por:</b> </p>
      <q-tabs class="q-ma-none q-pa-none text-purple">
        <q-route-tab icon="6_ft_apart" to="/adm/viajes/viajesCercanos" label="Distancia" exact />
        <q-route-tab icon="person" to="/adm/viajes/viajesByUsuario" label="Usuario" exact />
        <q-route-tab icon="people_alt" to="/adm/viajes/viajesByPasajero" label="Pasajero" exact />
        <q-route-tab icon="manage_accounts" to="/adm/viajes/viajesByConductor" label="Conductor" exact />
        <q-route-tab icon="airport_shuttle" to="/adm/viajes/viajesByCategoria" label="Categoria" exact />
      </q-tabs>
    </div>

    <ByFecha/>
  </div>
</template>

<script lang="ts">
import { pageStore } from '../../../stores/pageStore';
import ByFecha from './viajesByFecha.vue'
export default {
  name: 'IndexEmpresas',
  components:{ByFecha},
  setup() {
    const store = pageStore();

    return {
      store,
    };
  },
};
</script>
