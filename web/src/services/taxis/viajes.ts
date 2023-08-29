/* eslint-disable @typescript-eslint/no-explicit-any */
import { pageStore } from '../../stores/pageStore';
import gql from 'graphql-tag';
import { CreateViajes } from '../../types/taxis/viajes';

export default class ViajesService {
  async viajesCercanos(lat:number, lon:number,radioMts:number) {
    const store = pageStore();
    const sql = gql`
      query viajesCercanos($lat:Float!,$lon:Float!,$radioMts:Int!){
        viajesCercanos(lat:$lat,lon:$lon,radioMts:$radioMts){
          id
          pasajero_id
          conductor_id
          estado
          descripcion
          categoria_id
          origen_lat
          origen_lon
          destino_lat
          destino_lon
          registrado
          pasajero_username
          conductor_username
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { lat:lat, lon:lon, radioMts:radioMts  },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async viajesByUsuario(usuario_id:string) {
    const store = pageStore();
    const sql = gql`
      query viajesByUsuario($usuario_id:ID!){
        viajesByUsuario(usuario_id:$usuario_id){
          id
          pasajero_id
          conductor_id
          estado
          descripcion
          categoria_id
          origen_lat
          origen_lon
          destino_lat
          destino_lon
          registrado
          pasajero_username
          conductor_username
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { usuario_id:usuario_id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async viajesByPasajero(usuario_id:string) {
    const store = pageStore();
    const sql = gql`
      query viajesByPasajero($usuario_id:ID!){
        viajesByPasajero(usuario_id:$usuario_id){
          id
          pasajero_id
          conductor_id
          estado
          descripcion
          categoria_id
          origen_lat
          origen_lon
          destino_lat
          destino_lon
          registrado
          pasajero_username
          conductor_username
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { usuario_id:usuario_id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async viajesByConductor(usuario_id:string) {
    const store = pageStore();
    const sql = gql`
      query viajesByConductor($usuario_id:ID!){
        viajesByConductor(usuario_id:$usuario_id){
          id
          pasajero_id
          conductor_id
          estado
          descripcion
          categoria_id
          origen_lat
          origen_lon
          destino_lat
          destino_lon
          registrado
          pasajero_username
          conductor_username
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { usuario_id:usuario_id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async viajesByCategoria(categoria_id:string) {
    const store = pageStore();
    const sql = gql`
      query viajesByConductor($categoria_id:ID!){
        viajesByCategoria(categoria_id:$categoria_id){
          id
          pasajero_id
          conductor_id
          estado
          descripcion
          categoria_id
          origen_lat
          origen_lon
          destino_lat
          destino_lon
          registrado
          pasajero_username
          conductor_username
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { categoria_id:categoria_id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async createViaje(input:CreateViajes) {
    const store = pageStore();
    const sql = gql`
      mutation createViaje($input:CreateViajes!){
        createViaje(input:$input){
          id
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { input:input },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }
  async finalizarViaje(id:string) {
    const store = pageStore();
    const sql = gql`
      mutation finalizarViaje($id:ID!){
        finalizarViaje(id:$id){id pasajero_id conductor_id estado}
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { id:id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }
}
