package entity

import (
    . "github.com/Danceiny/dict-service/common"
    . "github.com/Danceiny/dict-service/common/FastJson"
)

type BaseEntity struct {
    ID          NodeId `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
    CreatedTime uint   `gorm:"column:created_time" json:"createdTime"`
    UpdatedTime uint   `gorm:"column:updated_time" json:"updatedTime"`
    // 不叫DeletedAt，避开gorm的软删除
    DeletedTime uint        `gorm:"column:deleted_time" json:"deletedTime"`
    Attr        *JsonObject `gorm:"-" json:"attr"`
}

func (entity BaseEntity) GetAttr() *JsonObject {
    return (&entity).Attr
}

func (entity *BaseEntity) SetAttr(bytes []byte) {
    o := Parse(bytes)
    entity.Attr = &o
}

func (entity BaseEntity) GetBid() BID {
    return entity.ID
}

type Entity interface {
    GetBid() BID
    GetAttr() *JsonObject
    SetAttr(bytes []byte)
}
type BID = interface{}

type DynamicAttrPlugin interface {
    GetAttr() JsonObject
    SetAttr(bytes []byte)
}

//
// func (bid *BID) String() {
//
// }
