/* eslint-disable @typescript-eslint/no-explicit-any */
import { pageStore } from '../../stores/pageStore';
import gql from 'graphql-tag';
import { CreateCategoriaVehiculos } from '../../types/taxis/categoria_vehiculos';

export default class CategoriaVehiculosService {
  async categoria_vehiculos() {
    const store = pageStore();
    const sql = gql`
      query{
        categoria_vehiculos{
          id
          nombre
          descripcion
          vehiculos
        }
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async categoria_vehiculo(id:string) {
    const store = pageStore();
    const sql = gql`
      query categoria_vehiculo($id:ID!){
        categoria_vehiculo(id:$id){id nombre descripcion vehiculos}
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { id:id },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }

  async createUpdateCategoriaVehiculo(input:CreateCategoriaVehiculos) {
    const store = pageStore();
    const sql = gql`
      mutation createUpdateCategoriaVehiculo($input:CreateCategoriaVehiculos!){
        createUpdateCategoriaVehiculo(input:$input){id}
      }
    `;

    const res: any = await store.run_graphql(
      sql,
      { input:input },
      { showNotificacion: false, showNotificacionError: true }
    );
    return res;
  }
}
