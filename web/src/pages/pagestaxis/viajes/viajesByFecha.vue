<template>
  <div class="q-pa-md">

    <q-table title="Viajes disponibles" :loading="loading" :rows="viajes" :columns="columns"
      row-key="name" hide-pagination :rows-per-page-options="[0]" :filter="filter">

      <template v-slot:top-right>
        <q-input outlined dense v-model.date="query.fecha_inicio" type="datetime-local" label="latitud:" class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>

        <q-input outlined dense v-model.date="query.fecha_fin" type="datetime-local" label="longitud:" class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>

        <q-checkbox v-model="query.include_disponibles" label="disponibles" class="q-pr-md">
          <q-tooltip class="bg-purple">
            Incluir los viajes que estan fuera del rango de fechas, que no tienen conductor asignado y tampoco estan finalizados
          </q-tooltip>
        </q-checkbox>

        <q-btn-group square flat outline class="q-pr-md">
          <q-btn outline color="green" label="buscar" :disable="loading" :loading="loading" size="small" icon="add" @click="listar()" />
        </q-btn-group>

        <q-input outlined dense debounce="300" v-model="filter" placeholder="buscar ..." >
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
      </template>

      <template v-slot:body-cell-pasajero_username="props">
        <q-td :props="props">
          <router-link :to="'/adm/usuarios/ver?username='+props.row.pasajero_username" class="text-orange">
            {{ props.row.pasajero_username }}
          </router-link>
        </q-td>
      </template>
      <template v-slot:body-cell-conductor_username="props">
        <q-td :props="props">
          <router-link :to="'/adm/usuarios/ver?username='+props.row.conductor_username" class="text-orange">
            {{ props.row.conductor_username }}
          </router-link>
        </q-td>
      </template>

      <template v-slot:body-cell-estado="props">
        <q-td :props="props">
          {{ props.row.estado?'activo':'finalizado' }}
        </q-td>
      </template>

      <template v-slot:body-cell-origen_lat="props">
        <q-td :props="props">
          <p class="q-my-none"> {{ props.row.origen_lat }} </p>
          <p class="q-my-none"> {{ props.row.origen_lon }} </p>
        </q-td>
      </template>
      <template v-slot:body-cell-destino_lat="props">
        <q-td :props="props">
          <p class="q-my-none"> {{ props.row.destino_lat }} </p>
          <p class="q-my-none"> {{ props.row.destino_lon }} </p>
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
                <q-item clickable v-ripple
                  :to="'/adm/viajes/finalizar?id=' + props.row.id">
                  <q-item-section avatar>
                    <q-icon color="primary" name="edit" right />
                  </q-item-section>
                  <q-item-section>Finalizar</q-item-section>
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
import { onMounted, ref } from 'vue';
import ViajesServices from '../../../services/taxis/viajes';
import { pageStore } from '../../../stores/pageStore';
import { QueryInput, QueryViajes, Viajes } from '../../../types/taxis/viajes';

const columns = [
  { name: 'id', label: 'ID', field: 'id', sortable: true },
  { name: 'pasajero_username', label: 'pasajero', field: 'pasajero_username', sortable: true },
  { name: 'conductor_username', label: 'conductor', field: 'conductor_username', sortable: true },
  { name: 'estado', label: 'estado', field: 'estado', sortable: true },
  { name: 'descripcion', label: 'descripcion', field: 'descripcion', sortable: true },
  { name: 'origen_lat', label: 'origen', field: 'origen_lat', sortable: true },
  { name: 'destino_lat', label: 'destino', field: 'destino_lat', sortable: true },
  { name: 'registrado', label: 'Solicitado', field: 'registrado', sortable: true },
  { name: 'opt', label: '', field: 'opt' },
];

export default {
  name: 'IndexEmpresas',
  setup() {
    const store = pageStore();
    const viajes = ref<Viajes[]>([]);
    const loading = ref(false);
    const viajesServices = new ViajesServices();
    const query = ref<QueryViajes>(QueryInput);

    async function listar(){
      viajes.value = [];
      loading.value = true;
      const res = await viajesServices.viajesByFecha(query.value).then((e) => e).catch((e) => e);
      viajes.value = res.viajesByFecha;
      loading.value = false;
    }

    onMounted(()=>{
      const gmtOffsetMillisGMT_4 = 240 * 60 * 1000;
      const now = new Date();
      const ahora = new Date(now.getTime() - gmtOffsetMillisGMT_4);
      ahora.setHours(-4);
      ahora.setMinutes(0);
      const dia_1 = new Date(ahora.getTime());
      dia_1.setHours(19);
      dia_1.setMinutes(59);
      const formattedDateTime = ahora.toISOString().slice(0, 16);
      const formattedDateTime2 = dia_1.toISOString().slice(0, 16);
      query.value.fecha_inicio = formattedDateTime;
      query.value.fecha_fin = formattedDateTime2;
      listar();
    })
    return {
      store,
      columns,
      viajes,
      loading,
      query,
      filter: ref(''),
      listar,
    };
  },
};
</script>
