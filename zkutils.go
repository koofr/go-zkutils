package zkutils

import (
	zk "launchpad.net/gozk"
	"strings"
)

func BuildPath(pathParts []string) string {
	return "/" + strings.Join(pathParts, "/")
}

func EnsurePath(z *zk.Conn, pathParts []string, perm []zk.ACL) (err error) {
	path := BuildPath(pathParts)

	stat, err := z.Exists(path)

	if err != nil || stat != nil {
		return
	}

	currentParts := make([]string, 0, len(pathParts))

	for _, part := range pathParts {
		currentParts = append(currentParts, part)

		_, err = z.Create(BuildPath(currentParts), "", 0, perm)

		if err != nil {
			if zkErr, ok := err.(*zk.Error); ok && zkErr.Code == zk.ZNODEEXISTS {
				err = nil
			} else {
				return
			}
		}
	}

	return
}
