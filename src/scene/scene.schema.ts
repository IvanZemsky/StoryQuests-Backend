import { Prop, Schema, SchemaFactory } from "@nestjs/mongoose"
import { ApiProperty } from "@nestjs/swagger"
import mongoose, { HydratedDocument, ObjectId } from "mongoose"
import { Answer } from "./types/types"
import { answersExample } from "./docs/docs"
import { Story } from "src/story/schemas/story.schema"

@Schema({ collection: "scenes", versionKey: false })
export class Scene {
   @ApiProperty({
      description: "Уникальный идентификатор сцены в базе данных",
      example: "66d0d0c6a63f85b540d92934",
   })
   _id: ObjectId

   @ApiProperty({
      description: "Уникальный идентификатор сцены в истории",
      example: "scene_1",
   })
   @Prop()
   sceneId: string

   @ApiProperty({
      description: "Уникальный идентификатор истории, к которой принадлежит сцена",
      example: "66cb6fb8ebae2e4b8fffd190",
   })
   @Prop({ type: mongoose.Schema.Types.ObjectId })
   storyId: Story

   @ApiProperty({
      description: "Заголовок сцены",
      example: "Dark Corridor",
   })
   @Prop()
   title: string

   @ApiProperty({
      description: "Описание сцены",
      example:
         "The doors creak open, revealing a dark corridor leading into the depths of the temple.",
   })
   @Prop()
   description: string

   @ApiProperty({
      description: "Тип сцены",
      example: "default",
   })
   @Prop()
   type: SceneType

   @ApiProperty({
      description: "Ссылка на иллюстрацию",
      example: "https://images.unsplash.com/photo-123",
   })
   @Prop()
   img: string

   @ApiProperty({
      description: "Ссылка на иллюстрацию",
      example: answersExample,
   })
   @Prop()
   answers: Answer[]

   @ApiProperty({
      description: "Количество прохождений для сцены с типом end",
      example: 99,
   })
   @Prop()
   passes?: number

   @Prop({ required: true })
   number: string
}

export type SceneType = "default" | "end"

export type SceneDocument = HydratedDocument<Scene>
export const SceneSchema = SchemaFactory.createForClass(Scene)
