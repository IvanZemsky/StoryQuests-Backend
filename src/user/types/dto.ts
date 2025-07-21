import { ApiProperty } from "@nestjs/swagger"

export class CreateUserDto {
   @ApiProperty({ example: "Curry", description: "Login" })
   readonly login: string

   @ApiProperty({ example: "123456", description: "Password" })
   readonly password: string
}