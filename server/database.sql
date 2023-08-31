create table `usuarios`(
    `id` integer unsigned not null auto_increment primary key, 
    `nombres` varchar(30) not null,
    `apellidos` varchar(30) not null,
    `username` varchar(25) not null unique,
    `password` varchar(100) not null,
    `telefono` varchar(20) not null,
    `documento` varchar(30) not null,
    `correo` varchar(100),
    `foto_url` varchar(200),
    `fecha_nac` date,
    `domicilio` varchar(100),
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    `estado` tinyint(1) not null default 1
);

create table `roles`(
    `id` integer unsigned not null auto_increment primary key, 
    `nombre` varchar(30) not null,
    `bit` integer unsigned not null, -- serial incremental
    CONSTRAINT `unique_rol_app` UNIQUE (`nombre`, `bit`)
);

create table `rol_usuario`(
    `usuario_id` integer unsigned not null,
    `rol_id` integer unsigned not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`usuario_id`) references `usuarios`(`id`),
    foreign key(`rol_id`) references `roles`(`id`),
    primary key(`usuario_id`, `rol_id`)
);

create table `rol_permiso`(
    `metodo` varchar(30) not null,
    `descripcion` varchar(200) not null,
    `rol_bits` integer unsigned not null, 
    primary key(`metodo`,`rol_bits`)
);

create table `tokens`(
    `username` varchar(25) not null,
    `token` varchar(500) not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP, 
    primary key(`username`)
);

-- =====================================================================================================================
-- =====================================================================================================================
-- =====================================================================================================================

create table `categorias`(
    `id` integer unsigned not null auto_increment primary key, 
    `nombre` varchar(100) not null,
    `descripcion` varchar(250) not null,
    `estado` tinyint(1) not null default 1,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP
);

create table `productos`(
    `id` integer unsigned not null auto_increment primary key, 
    `nombre` varchar(100) not null,
    `codigo` varchar(100) not null UNIQUE,
    `unidad_medida` varchar(50) not null,
    `descripcion` varchar(250) not null,
    `precio_entrada` DECIMAL(10, 2) not null,
    `precio_salida` DECIMAL(10, 2) not null,
    `categoria_id` integer unsigned not null, 
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP,
    foreign key(`categoria_id`) references `categorias`(`id`)
);

CREATE TABLE `personas` (
    `id` integer unsigned not null auto_increment primary key, 
    `nombres` VARCHAR(30) not null,
    `apellidos` VARCHAR(30) not null,
    `telefono` VARCHAR(20),
    `correo` VARCHAR(100),
    `direccion` VARCHAR(220),
    `tipo` ENUM('cliente', 'proveedor'),
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `transacciones` (
    `id` integer unsigned not null auto_increment primary key, 
    `tipo` ENUM('compra', 'venta'), 
    `descripcion` varchar(250) not null,
    `persona_id` integer unsigned not null,  
    `usuario_id` integer unsigned not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`persona_id`) REFERENCES personas(`id`),
    foreign key(`usuario_id`) references `usuarios`(`id`)
);  

CREATE TABLE `detalle_transacciones` (
    `id` integer unsigned not null auto_increment primary key, 
    `transaccion_id` integer unsigned not null,
    `producto_id` integer unsigned not null,
    `cantidad` integer not null,
    `precio_unitario` DECIMAL(10, 2) not null,
    FOREIGN KEY (`transaccion_id`) REFERENCES transacciones(`id`),
    FOREIGN KEY (`producto_id`) REFERENCES productos(`id`)
);

CREATE TABLE `almacenes` (
    `id` integer unsigned not null auto_increment primary key,
    `nombre` VARCHAR(100) not null,
    `descripcion` varchar(250) not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `ubicaciones` (
    `id` integer unsigned not null auto_increment primary key,
    `almacen_id` integer unsigned not null,
    `producto_id` integer unsigned not null,
    `cantidad` integer not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`almacen_id`) REFERENCES almacenes(`id`),
    FOREIGN KEY (`producto_id`) REFERENCES productos(`id`)
);


-- =====================================================================================================================
-- =====================================================================================================================
-- =====================================================================================================================


-- ALTER TABLE `usuarios` AUTO_INCREMENT = 1000;
-- ALTER TABLE `roles` AUTO_INCREMENT = 1000; 

insert into usuarios(nombres,apellidos,username,password,telefono,documento) values('Usuario','Administrador','admin','admin','+591 78227092','10721310');
insert into roles(nombre,bit) values('administrador',1);
insert into roles(nombre,bit) values('invitado',2);
insert into rol_usuario(usuario_id,rol_id) values(1,1);

-- administrador    = 1
-- invitado         = 2
-- pasajero         = 3
-- otros            = 4
-- se asigna la suma de valores de roles permitidos
-- valorAsignar = 1 << (rol_bit-1) 
-- se verifica: rol_bits & 4 == 4 para el rol 4
-- se verifica: 5 & 2 == 2 para el rol 2
insert into rol_permiso(rol_bits,metodo,descripcion) values
    (5,"usuarioByUsername","Recupera datos de un usuario por su username"),
    (1,"permisos","Recupera la lista de todos los permisos disponibles"),  
    (1,"roles","Recupera la lista de todos los roles disponibles"),  
    (1,"usuarios","Recupera la lista de todos los usuarios registrados"),  
    (1,"permisosByRol","Recupera la lista de todos los permisos disponibles por rol"),
    (1,"createUsuario","Registra un nuevo usuario"), 
    (1,"updateUsuario","Actualiza los datos de un usuario"), 
    (1,"createRol","Registra un nuevo Rol"),
    (1,"deleteRol","Elimina un rol del sistema"),
    (1,"modificarRol","Actualiza los datos de un rol"),
    (1,"usuariosByRol","Recupera la lista de todos los usuarios segun un rol");

-- ejercicios
-- agregar un usuario llamado juan con el rol de administrador
-- agregar un nuevo rol llamado asesor con los permisos de roles, permisos, usuarioByUsername
--  