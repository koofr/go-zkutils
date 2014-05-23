package zkutils_test

import (
	"github.com/koofr/go-netutils"
	. "github.com/koofr/go-zkutils"
	zk "github.com/koofr/gozk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
)

var _ = Describe("ZkUtils", func() {

	Describe("BuildPath", func() {
		It("should build path from parts", func() {
			Expect(BuildPath([]string{"part1", "part2"})).To(Equal("/part1/part2"))
		})
	})

	Describe("EnsurePath", func() {
		var s *TestServer
		var z *zk.Conn

		BeforeEach(func() {
			port, err := netutils.UnusedPort()
			Expect(err).NotTo(HaveOccurred())

			s, err = NewTestServer(port)
			Expect(err).NotTo(HaveOccurred())

			if err != nil {
				return
			}

			zz, session, err := zk.Dial("localhost:"+strconv.Itoa(port), 5e9)
			Expect(err).NotTo(HaveOccurred())

			Expect((<-session).State).To(Equal(zk.STATE_CONNECTED))

			z = zz
		})

		AfterEach(func() {
			z.Close()
			s.Stop()
		})

		It("should ensure path", func() {
			perm := zk.WorldACL(zk.PERM_ALL)

			stat, err := z.Exists("/part1")
			Expect(err).NotTo(HaveOccurred())
			Expect(stat).To(BeNil())

			err = EnsurePath(z, []string{"part1"}, perm)
			Expect(err).NotTo(HaveOccurred())

			stat, err = z.Exists("/part1")
			Expect(err).NotTo(HaveOccurred())
			Expect(stat).NotTo(BeNil())

			err = EnsurePath(z, []string{"part1"}, perm)
			Expect(err).NotTo(HaveOccurred())

			err = EnsurePath(z, []string{"part1", "part2", "part3"}, perm)
			Expect(err).NotTo(HaveOccurred())

			stat, err = z.Exists("/part1/part2")
			Expect(err).NotTo(HaveOccurred())
			Expect(stat).NotTo(BeNil())

			stat, err = z.Exists("/part1/part2/part3")
			Expect(err).NotTo(HaveOccurred())
			Expect(stat).NotTo(BeNil())
		})
	})

})
