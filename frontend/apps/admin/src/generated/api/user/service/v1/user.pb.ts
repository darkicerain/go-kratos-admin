// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.3
//   protoc               unknown
// source: user/service/v1/user.proto

/* eslint-disable */
import type { Empty } from "../../../google/protobuf/empty.pb";
import type { Timestamp } from "../../../google/protobuf/timestamp.pb";
import type { PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 用户权限 */
export enum UserAuthority {
  /** GUEST - 游客 */
  GUEST = "GUEST",
  /** CUSTOMER_USER - 普通用户 */
  CUSTOMER_USER = "CUSTOMER_USER",
  /** TENANT_ADMIN - 租户管理 */
  TENANT_ADMIN = "TENANT_ADMIN",
  /** SYS_ADMIN - 系统管理员 */
  SYS_ADMIN = "SYS_ADMIN",
}

/** 用户性别 */
export enum UserGender {
  /** SECRET - 未知 */
  SECRET = "SECRET",
  /** MALE - 男性 */
  MALE = "MALE",
  /** FEMALE - 女性 */
  FEMALE = "FEMALE",
}

/** 用户状态 */
export enum UserStatus {
  OFF = "OFF",
  ON = "ON",
}

/** 用户 */
export interface User {
  /** 用户ID */
  id?:
    | number
    | null
    | undefined;
  /** optional uint32 role_id = 2 [json_name = "roleId", (gnostic.openapi.v3.property) = {description: "角色ID"}];  // 角色ID */
  workId?:
    | number
    | null
    | undefined;
  /** 部门ID */
  orgId?:
    | number
    | null
    | undefined;
  /** 岗位ID */
  positionId?:
    | number
    | null
    | undefined;
  /** 租户ID */
  tenantId?:
    | number
    | null
    | undefined;
  /** 登录名 */
  username?:
    | string
    | null
    | undefined;
  /** 昵称 */
  nickname?:
    | string
    | null
    | undefined;
  /** 真实姓名 */
  realname?:
    | string
    | null
    | undefined;
  /** 头像 */
  avatar?:
    | string
    | null
    | undefined;
  /** 邮箱 */
  email?:
    | string
    | null
    | undefined;
  /** 手机号 */
  mobile?:
    | string
    | null
    | undefined;
  /** 座机号 */
  telephone?:
    | string
    | null
    | undefined;
  /** 性别 */
  gender?:
    | UserGender
    | null
    | undefined;
  /** 住址 */
  address?:
    | string
    | null
    | undefined;
  /** 国家地区 */
  region?:
    | string
    | null
    | undefined;
  /** 个人描述 */
  description?:
    | string
    | null
    | undefined;
  /** 备注名 */
  remark?:
    | string
    | null
    | undefined;
  /** 最后登录时间 */
  lastLoginTime?:
    | Timestamp
    | null
    | undefined;
  /** 最后登录IP */
  lastLoginIp?:
    | string
    | null
    | undefined;
  /** 用户状态 */
  status?:
    | UserStatus
    | null
    | undefined;
  /** 权限 */
  authority?:
    | UserAuthority
    | null
    | undefined;
  /** 角色码列表 */
  roles: string[];
  /** 创建者ID */
  createBy?:
    | number
    | null
    | undefined;
  /** 更新者ID */
  updateBy?:
    | number
    | null
    | undefined;
  /** 创建时间 */
  createTime?:
    | Timestamp
    | null
    | undefined;
  /** 更新时间 */
  updateTime?:
    | Timestamp
    | null
    | undefined;
  /** 删除时间 */
  deleteTime?: Timestamp | null | undefined;
}

/** 获取用户列表 - 答复 */
export interface ListUserResponse {
  items: User[];
  total: number;
}

/** 获取用户数据 - 请求 */
export interface GetUserRequest {
  id: number;
}

export interface GetUserByUserNameRequest {
  /** 用户登录名 */
  username: string;
}

/** 创建用户 - 请求 */
export interface CreateUserRequest {
  data:
    | User
    | null;
  /** 用户登录密码 */
  password?: string | null | undefined;
}

/** 更新用户 - 请求 */
export interface UpdateUserRequest {
  /** 用户的数据 */
  data:
    | User
    | null;
  /** 用户登录密码 */
  password?:
    | string
    | null
    | undefined;
  /** 要更新的字段列表 */
  updateMask:
    | string[]
    | null;
  /** 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。 */
  allowMissing?: boolean | null | undefined;
}

/** 删除用户 - 请求 */
export interface DeleteUserRequest {
  id: number;
}

/** 用户是否存在 - 请求 */
export interface UserExistsRequest {
  /** 用户登录名 */
  username: string;
}

/** 用户是否存在 - 答复 */
export interface UserExistsResponse {
  exist: boolean;
}

export interface BatchCreateUsersRequest {
  data: User[];
}

export interface BatchCreateUsersResponse {
  data: User[];
}

/** 用户服务 */
export interface UserService {
  /** 查询用户列表 */
  List(request: PagingRequest): Promise<ListUserResponse>;
  /** 查询用户详情 */
  Get(request: GetUserRequest): Promise<User>;
  /** 创建用户 */
  Create(request: CreateUserRequest): Promise<Empty>;
  /** 更新用户 */
  Update(request: UpdateUserRequest): Promise<Empty>;
  /** 删除用户 */
  Delete(request: DeleteUserRequest): Promise<Empty>;
  /** 批量创建租户 */
  BatchCreate(request: BatchCreateUsersRequest): Promise<BatchCreateUsersResponse>;
  /** 查询用户详情 */
  GetUserByUserName(request: GetUserByUserNameRequest): Promise<User>;
  /** 用户是否存在 */
  UserExists(request: UserExistsRequest): Promise<UserExistsResponse>;
}
