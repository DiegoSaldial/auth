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
