import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/LoginLayout.vue'),
    children: [],
  },

  {
    path: '/adm',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      {
        path: 'usuarios/',
        component: () => import('pages/usuarios/index.vue'),
      },
      {
        path: 'usuarios/ver',
        component: () => import('pages/usuarios/ver.vue'),
      },
      {
        path: 'usuarios/registrar',
        component: () => import('pages/usuarios/registrar.vue'),
      },
      {
        path: 'roles/',
        component: () => import('pages/roles/index.vue'),
      },
      {
        path: 'roles/permisos',
        component: () => import('pages/roles/permisos.vue'),
      },
      {
        path: 'roles/registrar',
        component: () => import('pages/roles/registrar.vue'),
      },
      {
        path: 'roles/eliminar',
        component: () => import('pages/roles/eliminar.vue'),
      },
      // ===============================================
      // ===============================================
      // ===============================================
      {
        path: 'categorias_vehiculos/',
        component: () => import('pages/pagestaxis/categoriasvehiculos/index.vue'),
      },
      {
        path: 'categorias_vehiculos/registrar',
        component: () => import('pages/pagestaxis/categoriasvehiculos/registrar.vue'),
      },
      {
        path: 'categorias_vehiculos/vehiculosByCategoria',
        component: () => import('pages/pagestaxis/categoriasvehiculos/vehiculosByCategoria.vue'),
      },
      {
        path: 'vehiculos/',
        component: () => import('pages/pagestaxis/vehiculos/index.vue'),
      },
      {
        path: 'vehiculos/registrar',
        component: () => import('pages/pagestaxis/vehiculos/registrar.vue'),
      },
      {
        path: 'vehiculos/asignarConductor',
        component: () => import('pages/pagestaxis/vehiculos/asignarConductor.vue'),
      },
      {
        path: 'vehiculos/quitarConductor',
        component: () => import('pages/pagestaxis/vehiculos/quitarConductor.vue'),
      },
      {
        path: 'conductores/',
        component: () => import('pages/pagestaxis/conductores/index.vue'),
      },
      {
        path: 'conductores/vehiculos',
        component: () => import('pages/pagestaxis/conductores/vehiculos.vue'),
      },
      {
        path: 'clientes/',
        component: () => import('pages/pagestaxis/clientes/index.vue'),
      },
      {
        path: 'viajes/',
        component: () => import('pages/pagestaxis/viajes/index.vue'),
      },
      {
        path: 'viajes/registrar',
        component: () => import('pages/pagestaxis/viajes/registrar.vue'),
      },
      {
        path: 'viajes/finalizar',
        component: () => import('pages/pagestaxis/viajes/finalizar.vue'),
      },
      {
        path: 'viajes/viajesCercanos',
        component: () => import('pages/pagestaxis/viajes/viajesCercanos.vue'),
      },
      {
        path: 'viajes/viajesByUsuario',
        component: () => import('pages/pagestaxis/viajes/viajesByUsuario.vue'),
      },
      {
        path: 'viajes/viajesByPasajero',
        component: () => import('pages/pagestaxis/viajes/viajesByPasajero.vue'),
      },
      {
        path: 'viajes/viajesByConductor',
        component: () => import('pages/pagestaxis/viajes/viajesByConductor.vue'),
      },
      {
        path: 'viajes/viajesByCategoria',
        component: () => import('pages/pagestaxis/viajes/viajesByCategoria.vue'),
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
