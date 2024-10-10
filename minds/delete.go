package minds

import (
	"fmt"
)

func (md *Minds) Delete(name string) error {

	// Delete the mind
	resp, err := md.Config.Delete("/projects/mindsdb/minds/" + name)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to delete mind: %s", name)
	}

	return nil

}
