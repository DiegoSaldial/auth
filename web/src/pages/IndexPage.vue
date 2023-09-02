
<template>
  <q-page class="items-center justify-evenly">
    <div class="row justify-center">
      <div class="col-xl-12 col-md-6 col-lg-4 q-mt-xl">
        <p class="text-center text-h5">Bienvenido al menu principal, aqui se visualizan los viajes disponibles</p>
      </div>
    </div>

    <div class="row justify-center">
      <div class="col-xs-12 col-sm-3">
        <q-input outlined dense v-model.date="query.fecha_inicio" type="datetime-local" label="Desde:" class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
      </div>
      <div class="col-xs-12 col-sm-3">
        <q-input outlined dense v-model.date="query.fecha_fin" type="datetime-local" label="Hasta:" class="q-mr-xl">
          <template v-slot:append>
            <q-icon name="search" />
          </template>
        </q-input>
      </div>
      <div class="col-xs-12 col-sm-3">
        <q-btn-group square flat outline class="q-pr-md">
          <q-btn outline color="green" :disable="loading" :loading="loading" label="buscar" size="small" icon="add" @click="listar()" />
        </q-btn-group>
        <q-checkbox v-model="ver_destinos" label="destinos" class="q-pr-md">
          <q-tooltip class="bg-purple">
            Muestra los destinos de los pasajeros si es que esta disponible
          </q-tooltip>
        </q-checkbox>
      </div>
    </div>
    <!--  -->

    <div class="row justify-center q-gutter-none">
      <div class="col-xs-12 q-px-md">
        <p class="q-mb-none text-center text-green"> {{ viajes.length }} resultados encontrados </p>
        <div :class="$q.dark.isActive?'darkmapa':''">
          <l-map v-if="showMap" :zoom="zoom" :center="center" :options="mapOptions" style="height: 72vh" >
            <l-tile-layer :url="url" :attribution="attribution"/>
            <div v-for="(alquileres, index) of viajes" :key="index">
              <l-marker :lat-lng="[alquileres.origen_lat,alquileres.origen_lon]">
                <l-popup>
                  <div>
                    <p style="margin-top:0;"> <b>Pasajero:</b> {{ alquileres.pasajero_username }} </p>
                    <p style="margin-top:0;margin-bottom:0"> Solicitado {{ store.parseDate(alquileres.registrado) }} </p>
                    <p style="margin-top:0;margin-bottom:0"> {{ alquileres.descripcion }} </p>
                  </div>
                </l-popup>
              </l-marker>

              <l-circle-marker color="red" v-if="ver_destinos && alquileres.destino_lat" :lat-lng="[alquileres.destino_lat,alquileres.destino_lon]">
                <l-popup>
                  <div>
                    <p class="q-my-none"> Este es el destino de {{ alquileres.pasajero_username }} </p>
                    <p class="q-my-none"> {{ alquileres.descripcion }} </p>
                  </div>
                </l-popup>
              </l-circle-marker>
            </div>
          </l-map>
        </div>
      </div>
    </div>
  </q-page>
</template>

<style>
  .darkmapa .leaflet-layer,
  .darkmapa .leaflet-control-zoom-in,
  .darkmapa .leaflet-control-zoom-out,
  .darkmapa .leaflet-control-attribution {
    filter: invert(100%) hue-rotate(180deg) brightness(95%) contrast(90%);
  }
</style>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import { QueryInput, QueryViajes, Viajes } from '../types/taxis/viajes';
import ViajesServices from '../services/taxis/viajes';
import 'leaflet/dist/leaflet.css';
import { latLng } from 'leaflet';
import { LMap, LTileLayer,LPopup,LMarker, LCircleMarker } from '@vue-leaflet/vue-leaflet';
import { pageStore } from '../stores/pageStore';

export default defineComponent({
  name: 'IndexPage',
  components: {
    LMap,
    LTileLayer,
    LPopup,
    LMarker,
    LCircleMarker,
  },
  setup() {
    const store = pageStore();
    const refmap = ref();
    const loading = ref(false);
    const query = ref<QueryViajes>(QueryInput);
    const viajes = ref<Viajes[]>([]);
    const viajesServices = new ViajesServices();
    const ver_destinos = ref(false);

    async function listar(){
      viajes.value = [];
      loading.value = true;
      query.value.include_disponibles = true;
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
      dia_1.setDate(dia_1.getDate()+1)
      dia_1.setHours(0);
      dia_1.setMinutes(59);
      const formattedDateTime = ahora.toISOString().slice(0, 16);
      const formattedDateTime2 = dia_1.toISOString().slice(0, 16);
      query.value.fecha_inicio = formattedDateTime;
      query.value.fecha_fin = formattedDateTime2;
      listar();
    })

    return {
      refmap,
      zoom: 14.2,
      center: latLng(-21.528098, -64.730105),
      url: 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',
      attribution:'&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
      showMap: true,
      mapOptions: {
        zoomSnap: 0.5
      },
      store,
      query,
      viajes,
      loading,
      ver_destinos,
      listar,
    };
  }
});
</script>
