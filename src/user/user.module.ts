import { forwardRef, Module } from "@nestjs/common"
import { UserController } from "./user.controller"
import { UserService } from "./user.service"
import { MongooseModule } from "@nestjs/mongoose"
import { User, UserSchema } from "./user.schema"
import { AuthModule } from "src/auth/auth.module"
import { StoryModule } from "src/story/story.module"

@Module({
   imports: [
      MongooseModule.forFeature([{ name: User.name, schema: UserSchema }]),
      forwardRef(() => AuthModule),
      forwardRef(() => StoryModule),
   ],
   controllers: [UserController],
   providers: [UserService],
   exports: [UserService],
})
export class UserModule {}
