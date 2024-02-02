package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("qricxmco574dent")
		if err != nil {
			return err
		}

		// add
		new_template := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "myj6hh43",
			"name": "template",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "ka687zbfn1u4iha",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_template)
		collection.Schema.AddField(new_template)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("qricxmco574dent")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("myj6hh43")

		return dao.SaveCollection(collection)
	})
}
